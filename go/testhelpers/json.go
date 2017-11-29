// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package testhelpers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"reflect"
)

type NormalizedJson map[string]interface{}

func GetNormalizedJson(a interface{}) (v NormalizedJson, err error) {
	if b, ok := a.([]byte); ok {
		err = json.Unmarshal(b, &v)
		return
	}

	if s, ok := a.(string); ok {
		err = json.Unmarshal([]byte(s), &v)
		return
	}

	b, err := json.Marshal(a)
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(b, &v)
	return
}

// IsJsonSubset returns true if |a| is a subset of |b|.
func IsJsonSubset(a, b interface{}) (bool, error) {
	av, err := GetNormalizedJson(a)
	if err != nil {
		return false, err
	}

	bv, err := GetNormalizedJson(b)
	if err != nil {
		return false, err
	}

	return isJsonSubsetMap(av, bv)
}

func isJsonSubsetMap(av, bv NormalizedJson) (bool, error) {
	for k, want := range av {
		has, ok := bv[k]
		if !ok {
			return false, errors.Errorf("required field '%s' not found", k)
		}

		if subset, err := isJsonSubsetInterface(want, has); !subset {
			return false, errors.Wrapf(err, "field '%s'", k)
		}
	}

	return true, nil
}

func isJsonSubsetInterface(want, has interface{}) (bool, error) {
	switch wt := want.(type) {
	case NormalizedJson:
		if subset, err := isJsonSubsetMap(wt, has.(NormalizedJson)); !subset {
			return false, err
		}

	case bool:
		if !wt {
			return true, nil
		}
		ht := has.(bool)
		if !ht {
			return false, errors.New("field is not equal")
		}

	case []interface{}:
		ht := has.([]interface{})
		if len(wt) != len(ht) {
			return false, errors.Errorf("array length mismatch. expected %d, found %d", len(wt), len(ht))
		}

		for i := range wt {
			if subset, err := isJsonSubsetInterface(wt[i], ht[i]); !subset {
				return false, errors.Wrapf(err, "index %d", err)
			}
		}

	default:
		wv := reflect.ValueOf(wt)
		hv := reflect.ValueOf(has).Convert(wv.Type())
		if wv.Type() != hv.Type() {
			return false, errors.Errorf("type mismatch. expected %s, found %s", wv.Type().String(), hv.Type().String())
		}

		if !hv.Type().Comparable() {
			return false, errors.Errorf("type %s is not comparable", hv.Type().String())
		}

		if !reflect.DeepEqual(want, has) {
			return false, errors.New("field was not deep-equal")
		}
	}

	return true, nil
}
