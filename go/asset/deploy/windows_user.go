// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"crypto/rand"
	"strings"

	"chromium.googlesource.com/enterprise/cel/go/common"
	assetpb "chromium.googlesource.com/enterprise/cel/go/schema/asset"
	commonpb "chromium.googlesource.com/enterprise/cel/go/schema/common"
	"github.com/pkg/errors"
)

type windowsUser struct{}

func (*windowsUser) ResolveGeneratedContent(ctx common.Context, u *assetpb.WindowsUser) (err error) {
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

	s := &commonpb.Secret{Final: []byte(p)}
	return ctx.Publish(u, "password", s)
}

// Characters that can appear in passwords.
const pwChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&'()*+,-./:;<=>?@[]^_`{|}~"

func generatePassword() (string, error) {
	entropy := make([]byte, 32)

	c, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}

	if c != len(entropy) {
		return "", errors.Errorf("Unexpected failure reading random bytes. Got length %d. Want %d", c, len(entropy))
	}

	// Don't start with a hyphen or a slash (will be interpreted as a flag)
	allowedPwChars := strings.Replace(pwChars, "-", "", 1)
	allowedPwChars = strings.Replace(allowedPwChars, "/", "", 1)

	pwd := ""
	for _, b := range entropy {
		idx := int(b) % len(allowedPwChars)
		pwd = pwd + string(allowedPwChars[idx])
		allowedPwChars = pwChars
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
		return errors.Errorf(`password cannot start with dash`)
	}

	// Strings starting with slash will be interpreted as flags by `net user`
	if strings.HasPrefix(password, "/") {
		return errors.Errorf(`password cannot start with slash`)
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
