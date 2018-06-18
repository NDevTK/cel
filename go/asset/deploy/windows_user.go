// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/asset"
	"chromium.googlesource.com/enterprise/cel/go/common"
	"crypto/rand"
	"github.com/pkg/errors"
)

type windowsUser struct{}

func (*windowsUser) ResolveGeneratedContent(ctx common.Context, u *asset.WindowsUser) (err error) {
	var p string
	if u.HardcodedPassword != "" {
		p = u.HardcodedPassword
	} else {
		p, err = generatePassword()
		if err != nil {
			return err
		}
	}

	s := &common.Secret{Final: []byte(p)}
	return ctx.Publish(u, "password", s)
}

// The list of characters that do not need escaping, i.e. they are passed as-is by cmd.exe
// as arguments.
const pwChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$'()*+,./:;=?@[\\]_`{}~"

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

func init() {
	common.RegisterResolverClass(&windowsUser{})
}
