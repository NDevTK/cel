// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/logging/logadmin"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
)

// Stay below project-wide quotas (per second - enforced per minute)
// See: https://cloud.google.com/logging/quotas
const waitTimeBetweenLogReads = 3 * time.Second

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
//    log entries labeled with the command id.
//
// 4. Meanwhile the calling code polls the logs looking for the command id
//    label. When one shows up with Labels["type"] == "result", it decodes the
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

	return operation.watchLogsForResult()
}

type runCommandOperation struct {
	ctx        context.Context
	project    string
	zone       string
	instance   string
	runCommand *RunCommandMetadataEntry

	lastLogInsertId string
	timeoutAt       time.Time
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

// There's sometime a noticeable (~2-5 seconds) lag between when a trace
// shows up in the Stackdriver vs Console logs. However, the Stackdriver
// Operation/Labels features helps us reliably distinguish command logs from
// other console logs, so we'll live with that extra cost for now.
// TODO: Improve command output printing latency
func (r *runCommandOperation) watchLogsForResult() (int, error) {
	r.resetTimeout()
	for r.timedout() {
		client, err := logadmin.NewClient(r.ctx, r.project)
		if err != nil {
			return -1, fmt.Errorf("can't read logs for project %s", r.project)
		}

		//todo: don't fetch what you just fetched (use timestamp? might still need a check for equality ...) //tODO: put project in there below. //TODO: instance filter.
		filter := fmt.Sprintf(`logName = "projects/%s/logs/%s"`, r.project, "cel%2Fcommander")
		if r.lastLogInsertId != "" {
			filter = filter + fmt.Sprintf(` AND insertId > "%s"`, r.lastLogInsertId)
		}
		filter = filter + fmt.Sprintf(` AND operation.id = "%s"`, r.runCommand.Id)
		entries := client.Entries(r.ctx, logadmin.Filter(filter))
		if err != nil {
			return -1, err
		}

		result, err := r.searchForCommandResult(entries)
		if err != nil {
			return -1, err
		}

		if result != nil {
			return result.ExitCode, nil
		}

		time.Sleep(waitTimeBetweenLogReads)
	}
	return -1, fmt.Errorf("timed out while waiting for the command to finish (%s)", r.runCommand.Id)
}

// This prints command output and looks for the special result log.
func (r *runCommandOperation) searchForCommandResult(entries *logadmin.EntryIterator) (*RunCommandResultLogEntry, error) {
	for {
		entry, err := entries.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to iterate through logs: %v", err)
		}

		r.lastLogInsertId = entry.InsertID
		r.resetTimeout()

		line := fmt.Sprintf("%v", entry.Payload)
		if v, ok := entry.Labels["type"]; ok && v == "result" {
			var result RunCommandResultLogEntry
			return &result, json.Unmarshal([]byte(line), &result)
		}

		// Don't print `result` - it's not real command output.
		fmt.Println(line)
	}

	return nil, nil
}

func (r *runCommandOperation) resetTimeout() {
	r.timeoutAt = time.Now().Add(timeoutNoSignal)
}

func (r *runCommandOperation) timedout() bool {
	return time.Now().Before(r.timeoutAt)
}
