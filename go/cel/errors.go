// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cel

import (
	"fmt"
)

// ConfigurationError is returned when a configuration file failed to load.
type ConfigurationError struct {
	// The file which resulted in the error.
	Filename string

	// True if the file was already loaded. In this case there would be no
	// underlying error.
	WasAlreadyLoaded bool

	// The underlying error, if there was one.
	UnderlyingError error
}

// Cause implements the Causer interface.
func (c *ConfigurationError) Cause() error {
	if c.UnderlyingError != nil {
		return c.UnderlyingError
	}
	return c
}

func (c *ConfigurationError) Error() string {
	if c.WasAlreadyLoaded {
		return fmt.Sprintf("the configuartion file %s was already loaded once", c.Filename)
	}

	return fmt.Sprintf("the configration file at %s failed to load: %s", c.Filename, c.UnderlyingError.Error())
}

func NewConfigurationError(filename string, already_loaded bool, underlying_error error) *ConfigurationError {
	return &ConfigurationError{filename, already_loaded, underlying_error}
}
