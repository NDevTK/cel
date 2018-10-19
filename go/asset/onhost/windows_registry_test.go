// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"testing"
)

func TestWindowsRegistryResolver_getRegistryValueTypeAndData(t *testing.T) {
	registry_value := &asset.RegistryValue{Name: "DwordValue", ValueType: &asset.RegistryValue_DwordValue{DwordValue: 123}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_DWORD", "123")

	registry_value = &asset.RegistryValue{Name: "DwordValueZeroIsNotUndefined", ValueType: &asset.RegistryValue_DwordValue{DwordValue: 0}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_DWORD", "0")

	registry_value = &asset.RegistryValue{Name: "QwordValue", ValueType: &asset.RegistryValue_QwordValue{QwordValue: 123}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_QWORD", "123")

	registry_value = &asset.RegistryValue{Name: "QwordValueZeroIsNotUndefined", ValueType: &asset.RegistryValue_QwordValue{QwordValue: 0}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_QWORD", "0")

	registry_value = &asset.RegistryValue{Name: "StringValue", ValueType: &asset.RegistryValue_StringValue{StringValue: "Some string"}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_SZ", "Some string")

	registry_value = &asset.RegistryValue{Name: "ExpandStringValue", ValueType: &asset.RegistryValue_ExpandStringValue{ExpandStringValue: "Some string"}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_EXPAND_SZ", "Some string")

	registry_value = &asset.RegistryValue{Name: "BinaryValue", ValueType: &asset.RegistryValue_BinaryValue{BinaryValue: []byte{0x46, 0x6f, 0x6f, 0x00, 0x01, 0x02}}}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_BINARY", "466f6f000102")

	multiStringValue := &asset.RegistryValue_MultiStringValue{MultiStringValue: &asset.RegistryValue_MultiString{Value: []string{"First", "Second", "Third"}}}
	registry_value = &asset.RegistryValue{Name: "MultiStringValueSeparatedByNulls", ValueType: multiStringValue}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_MULTI_SZ", "First\\0Second\\0Third")

	multiStringValue = &asset.RegistryValue_MultiStringValue{MultiStringValue: &asset.RegistryValue_MultiString{Value: []string{"First"}}}
	registry_value = &asset.RegistryValue{Name: "MultiStringValueSingleStringNoNull", ValueType: multiStringValue}
	testGetRegistryValueTypeAndData(t, registry_value, "REG_MULTI_SZ", "First")
}

func testGetRegistryValueTypeAndData(t *testing.T, registry_value *asset.RegistryValue, expected_type string, expected_data string) {
	r := &WindowsRegistryResolver{}

	value_type, value_data, err := r.GetRegistryValueTypeAndData(registry_value)

	if err != nil {
		t.Errorf("GetRegistryValueTypeAndData(%v) returned error %v", registry_value, err)
	}

	if value_type != expected_type {
		t.Errorf("GetRegistryValueTypeAndData(%v).type = %v, want %v", registry_value, value_type, expected_type)
	}

	if value_data != expected_data {
		t.Errorf("GetRegistryValueTypeAndData(%v).data = %v, want %v", registry_value, value_data, expected_data)
	}
}
