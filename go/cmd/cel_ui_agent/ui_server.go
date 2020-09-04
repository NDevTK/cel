// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

// The code does not handle concurrent calls to the methods of uiServer, since
// the only caller is run_ui_test.py, and it is guaranteed that the method calls
// are serialized.

// We explicitly processes timeout, instead of using exec.CommandContext()
// plus context.WithTimeout(). The reason is that because of this issue
// https://github.com/golang/go/issues/22381, the process cannot be reliably
// killed. Another reason is that we need to kill the chromedriver process
// explicitly. See the comment in timeOutHandler() for details.
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sync"
	"time"
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

// max time out is 1 hour
const maxTimeout int = 3600

// default time out is 5 minutes
const defaultTimeout int = 5 * 60

type runRequest struct {
	Command string
	Timeout int // time out in seconds.
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

	response := s.startCommand(request)

	json, err := json.Marshal(response)
	if err != nil {
		setInternalError(w)
		return
	}

	fmt.Fprint(w, string(json))
}

func (s *uiServer) startCommand(request runRequest) runResponse {
	log.Printf("request is: %#v", request)

	// validate time out
	if request.Timeout < 0 || request.Timeout > maxTimeout {
		return runResponse{
			Status:       statusError,
			ErrorMessage: fmt.Sprintf("Invalid Timeout value: %d. Valid range is (0, %d]", request.Timeout, maxTimeout),
		}
	}

	if request.Timeout == 0 {
		request.Timeout = defaultTimeout
	}

	s.command = exec.Command("powershell.exe", "-Command", request.Command)
	s.outputBuffer = &singleWriter{}
	s.command.Stdout = s.outputBuffer
	s.command.Stderr = s.outputBuffer
	err := s.command.Start()
	if err != nil {
		s.command = nil
		return runResponse{
			Status:       statusError,
			ErrorMessage: err.Error(),
		}
	}

	s.exitStatus = make(chan runStatus, 1)
	timer := time.AfterFunc(
		time.Duration(request.Timeout)*time.Second,
		s.timeoutHandler)
	go func() {
		s.waitForCommand()
		timer.Stop()
	}()

	return runResponse{
		Status: statusSuccess,
	}
}

// this method is called when timeout happens.
func (s *uiServer) timeoutHandler() {
	// time out happened
	log.Printf("Timeout")

	// kill the process
	err := exec.Command("taskkill", "/PID", string(rune(s.command.Process.Pid)), "/T", "/F").Run()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	// It's possible that the chromedriver process does not belong to the
	// process tree rooted at s.command.Process.Pid, thus, it will not be
	// killed by the previous command. So we kill it explicitly here
	// to ensures that chromedriver and the chrome processes started by
	// chromedriver all get killed.
	err = exec.Command("taskkill", "/IM", "chromedriver.exe", "/T", "/F").Run()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	// Also kill chrome, in case it is not started by chromedriver.
	err = exec.Command("taskkill", "/IM", "chrome.exe", "/T", "/F").Run()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	s.exitStatus <- runStatus{
		Status: statusKilled,
		Output: string(s.outputBuffer.b.Bytes()),
	}
}

func (s *uiServer) runStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	var response runStatus
	if s.command == nil {
		response = runStatus{
			Status: statusNoCommand,
		}
	} else {
		select {
		case status := <-s.exitStatus:
			response = status
			s.command = nil
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
	status := runStatus{}

	err := s.command.Wait()
	if err != nil {
		status.Status = statusFinishedWithError
	} else {
		status.Status = statusFinished
	}
	status.Output = string(s.outputBuffer.b.Bytes())
	s.exitStatus <- status
}
