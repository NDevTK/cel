// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"encoding/hex"
	"fmt"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"github.com/pkg/errors"
)

type WindowsRegistryResolver struct{}

func (w *WindowsRegistryResolver) ResolveOnHost(ctx common.Context, m *asset.WindowsMachine) error {
	d, ok := ctx.(*deployer)
	if !ok {
		return errors.New("ctx is not Deployer")
	}

	if m.Name == d.instanceName {
		for _, registry_key := range m.RegistryKey {
			if err := w.SetupRegistryKey(d, registry_key); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *WindowsRegistryResolver) SetupRegistryKey(d *deployer, registry_key *asset.RegistryKey) error {
	for _, value := range registry_key.Value {
		value_type, value_data, err := w.GetRegistryValueTypeAndData(value)

		if err != nil {
			return err
		}

		err = d.RunCommand("powershell.exe",
			"-File", d.GetSupportingFilePath("set_registry_value.ps1"),
			"-Path", registry_key.Path,
			"-Name", value.Name,
			"-Type", value_type,
			"-Data", value_data)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WindowsRegistryResolver) GetRegistryValueTypeAndData(value *asset.RegistryValue) (value_type string, value_data string, err error) {
	if x, ok := value.GetValueType().(*asset.RegistryValue_StringValue); ok {
		value_type = "REG_SZ"
		value_data = x.StringValue
	} else if x, ok := value.GetValueType().(*asset.RegistryValue_ExpandStringValue); ok {
		value_type = "REG_EXPAND_SZ"
		value_data = x.ExpandStringValue
	} else if x, ok := value.GetValueType().(*asset.RegistryValue_BinaryValue); ok {
		value_type = "REG_BINARY"
		value_data = hex.EncodeToString(x.BinaryValue)
	} else if x, ok := value.GetValueType().(*asset.RegistryValue_DwordValue); ok {
		value_type = "REG_DWORD"
		value_data = fmt.Sprintf("%d", x.DwordValue)
	} else if x, ok := value.GetValueType().(*asset.RegistryValue_QwordValue); ok {
		value_type = "REG_QWORD"
		value_data = fmt.Sprintf("%d", x.QwordValue)
	} else if x, ok := value.GetValueType().(*asset.RegistryValue_MultiStringValue); ok {
		value_type = "REG_MULTI_SZ"
		value_data = strings.Join(x.MultiStringValue.Value, "\\0")
	} else {
		return "", "", errors.Errorf("Couldn't find value type and data for registry_key: %s", value.Name)
	}

	return value_type, value_data, nil
}

func init() {
	common.RegisterResolverClass(&WindowsRegistryResolver{})
}
