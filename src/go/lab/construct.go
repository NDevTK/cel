// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package lab

func ConstructAssets(A *Assets, c *Config) (err error) {
	defer Action(&err, "constructing assets")

	err = ConstructProjectAsset(A, c)
	if err != nil {
		return
	}

	err = ConstructServiceAccountAssets(A, c)
	if err != nil {
		return
	}

	err = ConstructCryptoKeyAndPermissionAssets(A, c)
	if err != nil {
		return
	}

	err = ConstructNetworkAssets(A, c)
	if err != nil {
		return
	}

	err = ConstructInstanceAssets(A, c)
	if err != nil {
		return
	}

	return
}
