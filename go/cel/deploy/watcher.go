// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package deploy

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"fmt"
	runtimeconfig "google.golang.org/api/runtimeconfig/v1beta1"
	"strings"
	"time"
)

// Amount of time to wait for next round of polling for variables.
const pollDuration = time.Second * 30

// Print progress every Nth poll to prevent it from being too spammy.
const printProgressFrequency = 3

// Number of times to retry fetching variables before giving up.
// This prevents failures because of intermittent network issues.
const maxRetries = 3

// WaitForAllAssetsReady returns when assets are ready (or failed).
// We poll the runtime config variables every `interval` seconds to find out.
//
// TODO: Use variables.watch for low numbers of in-progress assets is small.
// TODO: Implement a second timeout value based on updateTime.
func WaitForAllAssetsReady(s *gcp.Session, timeout time.Duration) (err error) {
	defer common.LoggedAction(s.Logger, &err, "WaitForAllAssetsReady")()

	api, err := runtimeconfig.New(s.GetHttpClient())
	if err != nil {
		return err
	}

	configPath := fmt.Sprintf("projects/%s/configs/cel-config", s.GetProject())

	i := -1
	startTime := time.Now()
	for time.Now().Before(startTime.Add(timeout)) {
		variables, err := fetchVariablesFromRuntimeConfig(api, configPath)
		if err != nil {
			return err
		}

		assetStates, err := parseAssetStatesFromRuntimeConfigResponse(variables, configPath)
		if err != nil {
			return err
		}

		initAssets, hasInit := assetStates["init"]
		inProgressAssets, hasInProgress := assetStates["in-progress"]
		errorAssets, hasError := assetStates["error"]
		readyAssets, hasReady := assetStates["ready"]

		if hasError {
			return fmt.Errorf(`OnHost configuration failed on at least one asset: %v

See instance console logs for more info:
* https://console.cloud.google.com/compute/instances?project=%s`, errorAssets, s.GetProject())
		} else if hasInit || hasInProgress {
			// Only print progress once in a while so it's not too spammy
			if i++; i%printProgressFrequency == 0 {
				total := len(initAssets) + len(inProgressAssets)
				s.Logger.Info(common.MakeStringer("%d assets still deploying: %v", total, assetStates))
			}

			time.Sleep(pollDuration)
		} else if hasReady && len(assetStates) == 1 {
			s.Logger.Info(common.MakeStringer("All (%d) assets are ready: %s", len(readyAssets), readyAssets))
			return nil
		} else {
			s.Logger.Info(common.MakeStringer("Unknown asset states: %v", assetStates))
			return fmt.Errorf("failed to parse asset states")
		}
	}

	return fmt.Errorf(`OnHost configuration timed out

See instance console logs for more info:
* https://console.cloud.google.com/compute/instances?project=%s`, s.GetProject())
}

// Fetches the variables and values for this config.
func fetchVariablesFromRuntimeConfig(api *runtimeconfig.Service, configPath string) ([]*runtimeconfig.Variable, error) {
	retries := 0
	for {
		request := api.Projects.Configs.Variables.List(configPath).ReturnValues(true)
		response, err := request.Do()

		if err != nil {
			retries += 1
			if retries > maxRetries {
				return nil, err
			}
		}

		return response.Variables, nil
	}
}

// Parses runtime config variables into asset information.
// Returns a map of status to assets in this format:
//		{ status: [assetA, assetB, ...], ... }
//   with status=(init, in-progress, error, ready)
func parseAssetStatesFromRuntimeConfigResponse(variables []*runtimeconfig.Variable, configPath string) (map[string][]string, error) {
	assetStates := make(map[string][]string)
	for _, variable := range variables {
		name, status, err := parseAssetInfoFromVariable(variable, configPath)
		if err != nil {
			return nil, err
		}

		if assets, ok := assetStates[status]; ok {
			assetStates[status] = append(assets, name)
		} else {
			assetStates[status] = []string{name}
		}
	}

	return assetStates, nil
}

// Tries to parse asset information from a config variable.
func parseAssetInfoFromVariable(variable *runtimeconfig.Variable, configPath string) (name string, status string, err error) {
	name = variable.Name
	assetPrefix := fmt.Sprintf("%s/variables/asset/", configPath)
	statusSuffix := "/status"
	if !(strings.HasPrefix(name, assetPrefix) && strings.HasSuffix(name, statusSuffix)) {
		return "", "", fmt.Errorf("Couldn't parse asset state from '%v'", variable)
	}

	name = name[len(assetPrefix) : len(name)-len(statusSuffix)]
	status = variable.Text

	if variable.Text == "" {
		status = "init"
	}

	return name, status, nil
}
