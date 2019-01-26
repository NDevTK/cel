// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"github.com/pkg/errors"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

const defaultEtag = "NONE"

// The stamp file we use on start to know what's the last thing we processed.
var pathToLastCommandId = getWorkingDirectory() + "_LAST_PROCESSED_COMMAND_ID"

type commander struct {
	ctx           context.Context
	loggingClient *logging.Client
	logger        *logging.Logger
	deployer      *deployer
}

// The working directory where commands are executed with a trailing slash.
func getWorkingDirectory() string {
	if runtime.GOOS == "windows" {
		return "C:\\cel\\commander\\"
	} else {
		return "/cel/commander/"
	}
}

func CreateCommander(deployer *deployer) (*commander, error) {
	_ = os.Mkdir(getWorkingDirectory(), os.ModePerm)

	projectId, err := metadata.ProjectID()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx := context.Background()

	loggingClient, err := logging.NewClient(ctx, projectId)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	logger := loggingClient.Logger("cel/commander")

	return &commander{
		ctx:           ctx,
		loggingClient: loggingClient,
		logger:        logger,
		deployer:      deployer,
	}, nil
}

func (c *commander) Logf(format string, arg ...interface{}) {
	text := fmt.Sprintf(format, arg...)
	log.Output(2, text)
	c.logger.Log(
		logging.Entry{
			Payload: text,
		},
	)
}

func (c *commander) Close() error {
	return c.loggingClient.Close()
}

// Main loop for cel_agent's commander.
// Watches for RunCommand signals and calls ProcessRunCommandEntry.
func (c *commander) WatchForCommands() {
	c.Logf("Start to watch for commands.")
	defer c.Logf("Stopped watching for commands.")

	lastCommandId := seedLastProcessedCommandId()
	c.Logf("Set commander's LastCommandId=%v", lastCommandId)
	for {
		metadata, err := watchMetadata(c.ctx)
		if err != nil {
			c.Logf("Error during watchMetadata: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if c.ctx.Err() != nil {
			break
		}

		var runCommand gcp.RunCommandMetadataEntry
		json.Unmarshal([]byte(metadata.SerializedCommand), &runCommand)
		if lastCommandId != runCommand.Id {
			c.Logf("Detected a new RunCommandMetadataEntry: [runCommand.Id=%v, lastCommandId=%v]", runCommand.Id, lastCommandId)

			// Mark the command as processed now in case we crash or restart.
			// If this fails, it's not critical enough to abort because we
			// have `expireOn` as a backup to identify stale commands.
			_ = updateLastProcessedCommandId(runCommand.Id)
			lastCommandId = runCommand.Id

			// Launch asynchronously so we don't starve other commands.
			go c.ProcessRunCommandEntry(&runCommand)
		}

	}
}

// Execute the command we found in our metadata.
// Is responsible for logging the output and the exit code.
func (c *commander) ProcessRunCommandEntry(runCommand *gcp.RunCommandMetadataEntry) {
	// Skip very old commands - the caller probably timed out anyway.
	if runCommand.Expired() {
		c.Logf("Skipping RunCommand because it's expired: %v", runCommand)
		return
	}

	c.Logf("Processing command: %v", runCommand)

	operation := &logpb.LogEntryOperation{Id: runCommand.Id}
	exitCode := 0
	logInsertId := 0

	// Make our best effort to log `result` when we exit to release the caller.
	defer func() {
		result := &gcp.RunCommandResultLogEntry{ExitCode: exitCode}
		resultJson, err := json.Marshal(result)
		if err != nil {
			c.Logf("Error serializing results: %v", err)
		}

		resultStr := string(resultJson)
		log.Output(1, resultStr)
		c.logger.Log(
			logging.Entry{
				Payload:   resultStr,
				Operation: operation,
				Labels: map[string]string{
					"type": "result",
				},
				InsertID: formatLogInsertId(runCommand, logInsertId),
			},
		)

		// Ensure the `result` log is not buffered (client is waiting on this to return)
		c.logger.Flush()
	}()

	logFn := func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			// Text() usually returns a single line of output, but can also
			// return a partial line if it's over 64 * 1024 bytes.
			text := scanner.Text()

			log.Output(1, text)
			c.logger.Log(
				logging.Entry{
					Payload:   text,
					Operation: operation,
					InsertID:  formatLogInsertId(runCommand, logInsertId),
				},
			)
			logInsertId++
		}
	}

	// We assume that non-windows machines are Linux boxes hosting a NestedVM.
	// We make the same assumption in deployer.go.
	if runtime.GOOS == "windows" {
		exitCode = c.processRunCommandEntryOnWindows(runCommand, logFn)
	} else {
		exitCode = c.processRunCommandEntryOnNestedVM(runCommand, logFn)
	}
}

func (c *commander) processRunCommandEntryOnWindows(runCommand *gcp.RunCommandMetadataEntry, logFn func(io.Reader)) int {
	command := exec.Command("cmd.exe", "/C", runCommand.Command)

	return c.runCommand(command, logFn)
}

func (c *commander) runCommand(command *exec.Cmd, logFn func(io.Reader)) int {
	// Execute the command and redirect all output to our logger.
	command.Dir = getWorkingDirectory()

	var wg sync.WaitGroup
	wg.Add(1)
	stdout, err := command.StdoutPipe()
	command.Stderr = command.Stdout

	go func() {
		logFn(stdout)

		wg.Done()
	}()

	if err = command.Start(); err != nil {
		c.Logf("Error starting command: %v", err)
		return -1
	} else if err = command.Wait(); err != nil {
		c.Logf("Error waiting for command: %v", err)
		return -1
	}

	// Wait for all remaining output to be processed.
	wg.Wait()

	if !command.ProcessState.Success() {
		return -1
	}

	return 0
}

func (c *commander) processRunCommandEntryOnNestedVM(runCommand *gcp.RunCommandMetadataEntry, logFn func(io.Reader)) int {
	// gsutil is not installed on nested VMs, but it's a common command
	// used during tests. The host will act as a proxy to support gsutil.
	if strings.HasPrefix(runCommand.Command, "gsutil ") {
		return c.processGsutilOnNestedVM(runCommand, logFn)
	}

	output, err := c.deployer.RunCommandOnNestedVM("cmd.exe /C " + runCommand.Command)

	logFn(bytes.NewReader(output))

	if err != nil {
		return -1
	}

	return 0
}

func (c *commander) processGsutilOnNestedVM(runCommand *gcp.RunCommandMetadataEntry, logFn func(io.Reader)) int {
	// We currently only support "gsutil cp"
	re := regexp.MustCompile("^gsutil (cp|copy) (.*) (.*)$")
	match := re.FindStringSubmatch(runCommand.Command)
	if match == nil {
		logFn(strings.NewReader("gsutil is not installed on Nested VMs. Only `gsutil cp` is supported at this time."))

		return -1
	}

	cpDestLocal := getWorkingDirectory() + filepath.Base(match[2])
	cpDestVm := match[3]

	// Copy the file here (VM host)
	command := exec.Command("gsutil", match[1], match[2], cpDestLocal)
	exitCode := c.runCommand(command, logFn)
	if exitCode != 0 {
		return exitCode
	}

	// Upload the file to our nested VM
	err := c.deployer.UploadFileToNestedVM(cpDestLocal, cpDestVm)
	if err != nil {
		errorMessage := fmt.Sprintf("Error during UploadFileToNestedVM: %s", err)

		logFn(strings.NewReader(errorMessage))

		return -1
	}

	return 0
}

func formatLogInsertId(runCommand *gcp.RunCommandMetadataEntry, insertId int) string {
	return fmt.Sprintf("%s_%08d", runCommand.Id, insertId)
}

func seedLastProcessedCommandId() string {
	_, err := os.Stat(pathToLastCommandId)
	if err == nil {
		bytes, err := ioutil.ReadFile(pathToLastCommandId)
		if err != nil {
			return ""
		}
		return string(bytes)
	}
	return ""
}

func updateLastProcessedCommandId(commandId string) error {
	return ioutil.WriteFile(pathToLastCommandId, []byte(commandId), os.ModePerm)
}

// Code below is taken from GCEWindowsAgent's password reset feature:
// https://github.com/GoogleCloudPlatform/compute-image-windows
var (
	metadataURL    = "http://metadata.google.internal/computeMetadata/v1/instance/attributes"
	metadataHang   = "/?recursive=true&alt=json&wait_for_change=true&timeout_sec=60&last_etag="
	defaultTimeout = 70 * time.Second
	etag           = defaultEtag
)

func updateEtag(resp *http.Response) bool {
	oldEtag := etag
	etag = resp.Header.Get("etag")
	if etag == "" {
		etag = defaultEtag
	}
	return etag != oldEtag
}

type attributesJSON struct {
	// Wrap the serialized command inside the attributes structure.
	// We could query /instance/attributes/cel-command directly, but we won't
	// get an ETag to wait on when it doesn't exist (just a 404).
	SerializedCommand string `json:"cel-command"`
}

func watchMetadata(ctx context.Context) (*attributesJSON, error) {
	client := &http.Client{
		Timeout: defaultTimeout,
	}

	req, err := http.NewRequest("GET", metadataURL+metadataHang+etag, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Metadata-Flavor", "Google")
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	// Don't return error on a canceled context.
	if err != nil && ctx.Err() != nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// We return the response even if the etag has not been updated.
	updateEtag(resp)

	md, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var metadata attributesJSON
	return &metadata, json.Unmarshal(md, &metadata)
}
