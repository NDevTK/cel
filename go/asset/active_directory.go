// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package asset

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"crypto/rand"
	"github.com/pkg/errors"
)

func (w *WindowsUser) ResolveGeneratedContent(ctx common.Context) error {
	if w.Password.GetHardcoded() != "" {
		return ctx.Publish(w.Password, "final", []byte(w.Password.GetHardcoded()))
	}

	p, err := generatePassword()
	if err != nil {
		return err
	}

	return ctx.Publish(w.Password, "final", []byte(p))
}

// The printable characters from 0x21-0x7e rearranged to have a bias for
// alphanums. Due to the method of password generation, the first 64 characters
// have a 50% higher probability of being picked than others.
const pwChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

func generatePassword() (string, error) {
	entropy := make([]byte, 32)

	c, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}

	if c != len(entropy) {
		return "", errors.Errorf("Unexpected failure reading random bytes. Got length %d. Want %d", c, len(entropy))
	}

	pwd := ""
	for _, b := range entropy {
		idx := int(b) % len(pwChars)
		pwd = pwd + string(pwChars[idx])
	}

	return pwd, nil
}
