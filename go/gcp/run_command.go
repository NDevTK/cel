// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	compute "google.golang.org/api/compute/v1"
)

// How long to wait between GetSerialPortOutput calls. No known quota.
const waitTimeBetweenLogReads = 1 * time.Second

// How long before we timeout the operation when there're no new logs.
// This can be caused by cel_agent crashing or the command hanging.
// TODO: cel_agent should emit periodic type='processing' entries that are not
//       printed out, similar to what we do with type='result'.
const timeoutNoSignal = 2 * time.Minute

// RunCommand runs a command on a Windows VM hosted on Google Compute Engine.
//
// It uses a similar flow to GCE Windows Agent's password reset mechanism
// that we reimplemented in cel_agent.
// The code for the original password reset mechanism can be found at
// https://github.com/GoogleCloudPlatform/compute-image-windows.
//
// At a high level, RunCommand works as follows:
//
// 1. Someone with edit privileges over the GCE project adds a metadata entry
//    for an instance with key the 'cel-command', and JSON object as the value.
//    The value encodes a RunCommandMetadataEntry with information about what
//    should be executed.
//
// 2. The cel_agent running in the VM observes the metadata change, and picks
//    up the new 'cel-command' metadata entry.
//
// 3. The agent runs the command while logging all output and the exit code in
//    console output labeled with the command id.
//
// 4. Meanwhile the calling code polls the logs looking for our runcommand
//    label. When one shows up with a type == "result", it decodes the
//    payload and returns.
//
// The http.Client referenced by |client| will be used for interacting with the
// Google Cloud API and should already incorporate any required authentication.
//
// If the request is successful, the function prints the output of the command
// and returns the exit code. Otherwise, it will return an error.
func RunCommandOnInstance(ctx context.Context, client *http.Client,
	project, zone, instance string,
	runCommand *RunCommandMetadataEntry) (int, error) {
	operation := &runCommandOperation{
		ctx:        ctx,
		project:    project,
		zone:       zone,
		instance:   instance,
		runCommand: runCommand,
	}

	err := operation.setRunCommandMetadata(client)
	if err != nil {
		return -1, err
	}

	return operation.watchLogsForResult(client)
}

type runCommandOperation struct {
	ctx        context.Context
	project    string
	zone       string
	instance   string
	runCommand *RunCommandMetadataEntry

	lastOutputPosition int64
	timeoutAt          time.Time
	lastError          error
}

type logEntry struct {
	runCommandId string
	logType      string
	message      string
}

// This indicates to cel_agent that there is something to run on this instance.
// Instance.SetMetadata works with all items as a whole. This deals with that.
func (r *runCommandOperation) setRunCommandMetadata(client *http.Client) error {
	service, err := compute.New(client)
	if err != nil {
		return err
	}

	// Fetch the current Instance's Metadata and Status
	instance_data, err := service.Instances.Get(r.project, r.zone, r.instance).
		Context(r.ctx).Fields("status", "metadata").Do()
	if err != nil {
		return err
	}

	// Verify the target Compute Instance is running:
	// https://cloud.google.com/compute/docs/instances/checking-instance-status
	if instance_data.Status != "RUNNING" {
		return fmt.Errorf("expected status to be RUNNING, but found %s", instance_data.Status)
	}

	metadata := instance_data.Metadata

	// Set the cel-command metadata entry
	err = UpdateGceMetadataWithRunCommand(metadata, r.runCommand)
	if err != nil {
		return err
	}

	// Set the new Instance Metadata set
	_, err = service.Instances.SetMetadata(r.project, r.zone, r.instance, metadata).Context(r.ctx).Do()

	return err
}

// We previously used Stackdriver for its Operation/Labels features that helped
// us reliably distinguish command logs from other console logs, but it had a
// significant buffering delay (5+ seconds) and other weird behavior like out
// of order log visibility. We're now using Console logs with a special format
// that gives us better latency (<2s) and guarantees line order.
func (r *runCommandOperation) watchLogsForResult(client *http.Client) (int, error) {
	r.resetTimeout()
	for r.timedout() {
		service, err := compute.New(client)
		if err != nil {
			return -1, err
		}

		request := service.Instances.GetSerialPortOutput(r.project, r.zone, r.instance)
		if r.lastOutputPosition != 0 {
			request = request.Start(r.lastOutputPosition)
		}

		serialPortOutput, err := request.Context(r.ctx).Do()
		if err != nil {
			return -1, err
		}

		// Remove any partial lines from the console output.
		output := serialPortOutput.Contents
		lastNewline := strings.LastIndex(output, "\n")
		if lastNewline != -1 {
			lengthPartialLine := len(output) - lastNewline - 1
			if lengthPartialLine != 0 {
				output = output[:lastNewline+1]
			}

			// Update the start position for the next console fetch.
			r.lastOutputPosition = serialPortOutput.Next - int64(lengthPartialLine)

			// Parse, filter logs and print logs.
			result, err := r.searchForCommandResult(output)
			if err != nil {
				return -1, err
			}

			if result != nil {
				return result.ExitCode, nil
			}
		} else if len(output) > 512*1024 {
			// Our algorithm doesn't support lines longer than 512 KB because
			// of limitations on the GetSerialPortOutput side (1MB).
			return -1, fmt.Errorf("console output line bigger than 512KB (%d)", len(output))
		}

		time.Sleep(waitTimeBetweenLogReads)
	}
	return -1, fmt.Errorf("timed out while waiting for the command to finish (%s) (lastError=%v)", r.runCommand.Id, r.lastError)
}

// This prints command output and looks for the special result log.
func (r *runCommandOperation) searchForCommandResult(output string) (*RunCommandResultLogEntry, error) {
	for _, line := range strings.Split(output, "\n") {
		entry, err := parseLogEntryMessage(line)
		if err != nil {
			r.lastError = fmt.Errorf("failed to parse entry message [err=%v]", err)
			continue
		}

		if entry.runCommandId != r.runCommand.Id {
			r.lastError = fmt.Errorf("unexpected runCommandId [entry=%s][current=%s]", entry.runCommandId, r.runCommand.Id)
			continue
		}

		r.resetTimeout()

		if entry.logType == "result" {
			var result RunCommandResultLogEntry
			return &result, json.Unmarshal([]byte(entry.message), &result)
		}

		// Don't print `result` - it's not real command output.
		r.lastError = nil
		fmt.Println(entry.message)
	}

	return nil, nil
}

// Formats a single-line log message in a way that `run` clients can parse.
func FormatLogEntryMessage(runCommand *RunCommandMetadataEntry, logType string, message string) string {
	return fmt.Sprintf("[$RunCommand$%s$%s] %s", runCommand.Id, logType, message)
}

// Parses a single-line log message formatted by FormatLogEntryMessage.
var logEntryRegexp = regexp.MustCompile("\\[\\$RunCommand\\$([a-z0-9]+)\\$([a-z]*)\\] (.*)")

func parseLogEntryMessage(message string) (*logEntry, error) {
	// Quick check to see if this looks like a RunCommand log.
	if strings.Index(message, "[$RunCommand$") == -1 {
		return nil, fmt.Errorf("doesn't look like a RunCommand log entry: %s", message)
	}

	parts := logEntryRegexp.FindStringSubmatch(message)

	if len(parts) != 4 {
		return nil, fmt.Errorf("doesn't match our RunCommand log entry regex: %s", message)
	}

	return &logEntry{runCommandId: parts[1], logType: parts[2], message: parts[3]}, nil
}

func (r *runCommandOperation) resetTimeout() {
	r.timeoutAt = time.Now().Add(timeoutNoSignal)
}

func (r *runCommandOperation) timedout() bool {
	return time.Now().Before(r.timeoutAt)
}
