// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gcp

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	compute "google.golang.org/api/compute/v1"
	"time"
)

const (
	// Name of RunCommand metadata key.
	RunCommandKey = "cel-command"

	// Lifetime of 'cel-command' metadata key. If the agent sees the metadata
	// request outside of the lifetime, then that entry will be ignored.
	RunCommandLifetime = time.Minute * 5
)

type RunCommandMetadataEntry struct {
	// Unique identifier for the RunCommand entry.
	// This is used by cel_agent to dedupe recent commands after restarts.
	Id string `json:"id"`

	// The command to execute on the instance.
	Command string `json:"command"`

	// Time after which this metadata entry should be ignored. Should be
	// encoded as per RFC 3339.
	ExpireOn string `json:"expireOn"`
}

type RunCommandResultLogEntry struct {
	ExitCode int `json:"exit_code"`
}

func (e *RunCommandMetadataEntry) Expired() bool {
	t, err := time.Parse(time.RFC3339, e.ExpireOn)
	if err != nil {
		return true
	}
	return t.Before(time.Now())
}

func NewRunCommand(command string) *RunCommandMetadataEntry {
	expireOn := time.Now().Add(RunCommandLifetime).Format(time.RFC3339)

	h := sha1.New()
	h.Write([]byte(expireOn))
	h.Write([]byte(command))
	id := hex.EncodeToString(h.Sum(nil))

	return &RunCommandMetadataEntry{
		Id:       id,
		Command:  command,
		ExpireOn: expireOn}
}

func UpdateGceMetadataWithRunCommand(m *compute.Metadata, run_command *RunCommandMetadataEntry) error {
	value_set := false
	for _, item := range m.Items {
		if item.Key == RunCommandKey {
			value_buffer, err := json.Marshal(run_command)
			if err != nil {
				return err
			}
			item.Value = new(string)
			*item.Value = string(value_buffer)
			value_set = true
			break
		}
	}
	if !value_set {
		value_buffer, err := json.Marshal(run_command)
		if err != nil {
			return err
		}
		value := new(string)
		*value = string(value_buffer)
		m.Items = append(m.Items, &compute.MetadataItems{
			Key:   RunCommandKey,
			Value: value})
	}
	return nil
}
