// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package host

import (
	"github.com/pkg/errors"
)

func (*HostEnvironment) Validate() error { return nil }
func (*AddressPool) Validate() error     { return nil }
func (*MachineType) Validate() error     { return nil }
func (*Project) Validate() error         { return nil }
func (*Image_Family) Validate() error    { return nil }

func (i *Image) Validate() error {
	if i.GetSource() == nil {
		return errors.New("either 'latest' or 'url' must be specified")
	}
	return nil
}
