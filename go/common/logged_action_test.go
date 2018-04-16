// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type wrappedEvent struct {
	Message  string      `json:"m,omitempty"`
	Severity string      `json:"s"`
	Event    interface{} `json:"e,omitempty"`
}

type fakeLogger struct {
	events []string
}

func (f *fakeLogger) logit(v fmt.Stringer, severity string) {
	w := wrappedEvent{
		Message:  v.String(),
		Severity: severity,
		Event:    v,
	}

	b, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}

	s := string(b)
	f.events = append(f.events, s)
}

func (f *fakeLogger) Debug(v fmt.Stringer) {
	f.logit(v, "debug")
}

func (f *fakeLogger) Info(v fmt.Stringer) {
	f.logit(v, "info")
}

func (f *fakeLogger) Warning(v fmt.Stringer) {
	f.logit(v, "warning")
}

func (f *fakeLogger) Error(v fmt.Stringer) {
	f.logit(v, "error")
}

func TestLoggedAction_basic(t *testing.T) {
	l := &fakeLogger{}
	func() {
		var err error
		defer LoggedAction(l, &err, "hello %s", "world")()
	}()
	if len(l.events) != 2 ||
		l.events[0] != "{\"m\":\"[BEGIN] hello world\",\"s\":\"info\",\"e\":{\"m\":\"hello world\",\"s\":\"BEGIN\"}}" ||
		l.events[1] != "{\"m\":\"[  END] hello world\",\"s\":\"info\",\"e\":{\"m\":\"hello world\",\"s\":\"END\"}}" {
		t.Errorf("unexpected log contents: %#v", l.events)
	}
}

func TestLoggedAction_error(t *testing.T) {
	l := &fakeLogger{}
	func() (err error) {
		defer LoggedAction(l, &err, "hello %s", "world")()
		return errors.New("Oh no!")
	}()
	if len(l.events) != 2 ||
		l.events[0] != "{\"m\":\"[BEGIN] hello world\",\"s\":\"info\",\"e\":{\"m\":\"hello world\",\"s\":\"BEGIN\"}}" ||
		l.events[1] != "{\"m\":\"[ FAIL] hello world: Oh no!\",\"s\":\"error\",\"e\":{\"m\":\"hello world\",\"s\":\"FAIL\",\"err\":{}}}" {
		// Note that the error ended up looking like it's empty because it's
		// not JSON serializable. The error.Error() value is available via
		// event.String().
		t.Errorf("unexpected log contents: %#v", l.events)
	}
}

func TestLoggedAction_panic(t *testing.T) {
	l := &fakeLogger{}
	func() (err error) {
		defer LoggedAction(l, &err, "hello %s", "world")()
		panic(errors.New("Oh no!"))
	}()
	if len(l.events) != 2 ||
		l.events[0] != "{\"m\":\"[BEGIN] hello world\",\"s\":\"info\",\"e\":{\"m\":\"hello world\",\"s\":\"BEGIN\"}}" ||
		l.events[1] != "{\"m\":\"[ FAIL] hello world: panic: Oh no!\",\"s\":\"error\",\"e\":{\"m\":\"hello world\",\"s\":\"FAIL\",\"err\":{}}}" {
		// Note that the error ended up looking like it's empty because it's
		// not JSON serializable. The error.Error() value is available via
		// event.String().
		t.Errorf("unexpected log contents: %#v", l.events)
	}
}
