// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

// The code does not handle concurrent calls to the methods of uiServer, since
// the only caller is run_ui_test.py, and it is guaranteed that the method calls
// are serialized.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

type singleWriter struct {
	b  bytes.Buffer
	mu sync.Mutex
}

func (w *singleWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.b.Write(p)
}

type uiServer struct {
	command *exec.Cmd

	// the buffer to receive stdout and stderr of the commmand
	outputBuffer *singleWriter

	// This channel is used by waitForCommand() to transfer the
	// command status to runStatusHandler().
	exitStatus chan runStatus
}

type runRequest struct {
	Command string
	timeOut int // time out in seconds. -1 means infinite timeout
}

// RunStartStatus is the status of runStart request
type RunStartStatus int

const (
	// Command is succesfully started
	statusSuccess RunStartStatus = 0

	// Error occurred when starting the command
	statusError RunStartStatus = 1
)

type runResponse struct {
	Status       RunStartStatus
	ErrorMessage string
}

// RunStatus is the status of the command run
type RunStatus int

const (
	statusFinished          RunStatus = 0
	statusFinishedWithError RunStatus = 1
	statusRunning           RunStatus = 2
	statusKilled            RunStatus = 3 //  killed because of time out
	statusNoCommand         RunStatus = 4
)

type runStatus struct {
	Status RunStatus

	Output string
}

func (s *uiServer) start() {
	http.HandleFunc("/Run", s.runHandler)
	http.HandleFunc("/RunStatus/", s.runStatusHandler)
	port := "9000"

	// this service only accepts requests made from the localhost, to make
	// sure it cannot be used to gain unauthorized access to the instance.
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", port), nil)
}

func setBadRequestError(w http.ResponseWriter) {
	http.Error(w, "Bad Request", http.StatusBadRequest)
}

func setInternalError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (s *uiServer) runHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		setBadRequestError(w)
		return
	}

	var request runRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		setBadRequestError(w)
		return
	}

	log.Printf("command is: %s", request.Command)
	s.command = exec.Command("cmd.exe", "/C", request.Command)
	s.outputBuffer = &singleWriter{}
	s.command.Stdout = s.outputBuffer
	s.command.Stderr = s.outputBuffer
	err = s.command.Start()
	var response runResponse
	if err != nil {
		s.command = nil
		response = runResponse{
			Status:       statusError,
			ErrorMessage: err.Error(),
		}
	} else {
		response = runResponse{
			Status: statusSuccess,
		}
	}

	s.exitStatus = make(chan runStatus)
	go s.waitForCommand()

	json, err := json.Marshal(response)
	if err != nil {
		setInternalError(w)
		return
	}

	fmt.Fprint(w, string(json))
}

func (s *uiServer) runStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	var response runStatus
	if s.command == nil {
		response = runStatus{
			Status: 4,
		}
	} else {
		select {
		case status := <-s.exitStatus:
			response = status
		default:
			response = runStatus{
				Status: statusRunning,
				Output: "",
			}
		}
	}

	json, err := json.Marshal(response)
	if err != nil {
		setInternalError(w)
		return
	}
	fmt.Fprint(w, string(json))
}

// waitForCommand waits for the current command to finish
func (s *uiServer) waitForCommand() {
	err := s.command.Wait()
	status := runStatus{}
	if err != nil {
		status.Status = statusFinishedWithError
	} else {
		status.Status = statusFinished
	}
	status.Output = string(s.outputBuffer.b.Bytes())
	s.exitStatus <- status
}
