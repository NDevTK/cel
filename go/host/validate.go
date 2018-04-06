// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"github.com/pkg/errors"
	"strings"
)

func (*AddressPool) Validate() error     { return nil }
func (*HostEnvironment) Validate() error { return nil }
func (*Image_Family) Validate() error    { return nil }
func (*LogSettings) Validate() error     { return nil }
func (*MachineType) Validate() error     { return nil }
func (*Project) Validate() error         { return nil }
func (*RuntimeSupport) Validate() error  { return nil }

func (i *Image) Validate() error {
	if i.GetFixed() == "" && i.GetLatest() == nil {
		return errors.New("either 'url' or 'latest' must be specified for an Image")
	}
	return nil
}

func (s *Storage) Validate() error {
	if s.Prefix == "" {
		return nil
	}

	if strings.HasSuffix(s.Prefix, "/") {
		return errors.Errorf("the GCS object name prefix must not end with a slash. Found name \"%s\"", s.Prefix)
	}
	return nil
}
