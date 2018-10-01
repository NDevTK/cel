// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"crypto/rand"
	"github.com/pkg/errors"
	"strings"
)

type windowsUser struct{}

func (*windowsUser) ResolveGeneratedContent(ctx common.Context, u *asset.WindowsUser) (err error) {
	var p string
	if u.HardcodedPassword != "" {
		p = u.HardcodedPassword
		if err := validatePassword(p); err != nil {
			return err
		}
	} else {
		p, err = generatePassword()
		if err != nil {
			return err
		}
	}

	s := &common.Secret{Final: []byte(p)}
	return ctx.Publish(u, "password", s)
}

// The printable characters from 0x21-0x7e rearranged to have a bias for
// alphanums. Due to the method of password generation, the first 64 characters
// have a 50% higher probability of being picked than others.
const pwChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

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

// validatePassword validates that password is safe, i.e.
// - it only contains chars from pwChars, and
// - it can be passed as-is as arguments
func validatePassword(password string) error {
	if password == "" {
		return errors.Errorf("password cannot be empty")
	}

	// Strings starting with dash will be interpreted as flags by powershell
	if strings.HasPrefix(password, "-") {
		return errors.Errorf(`passowrd cannot start with dash`)
	}

	for _, ch := range password {
		if !strings.Contains(pwChars, string(ch)) {
			return errors.Errorf(`password cannot contain character "%s"`, string(ch))
		}
	}
	return nil
}

func init() {
	common.RegisterResolverClass(&windowsUser{})
}
