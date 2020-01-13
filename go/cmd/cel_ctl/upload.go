// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"chromium.googlesource.com/enterprise/cel/go/gcp"
	"github.com/spf13/cobra"
)

type UploadCommand struct {
	UseBuiltins bool
	Instance    string
	File        string
	Destination string
}

func init() {
	uc := &UploadCommand{}
	cmd := &cobra.Command{
		Use:   "upload [--instance machine] --file file [--destination dest_dir] [configuration files]",
		Short: "upload a file to a CELab instance",
		Long: `Uploads a file to a CELab instance from a lab previously deployed via a 'Deploy' command.
Will skip uploading to the Storage if the file already exists & has the same hash.
`}
	cmd.Flags().StringVar(&uc.Instance, "instance", "", "short instance name of VM to upload the file to (default: None)")
	cmd.Flags().StringVar(&uc.File, "file", "", "path to the file to upload")
	cmd.MarkFlagRequired("file")
	cmd.Flags().StringVar(&uc.Destination, "destination", ".", "destination directory of the upload (defaults to .)")
	cmd.Flags().BoolVarP(&uc.UseBuiltins, "builtins", "B", false, "Use builtin assets")
	app.AddCommand(cmd, uc)
}

func (uc *UploadCommand) Run(ctx context.Context, a *Application, cmd *cobra.Command, args []string) error {
	log.Printf("Start of `cel_ctl upload` - version %s", version)

	var err error
	defer func() {
		if err != nil {
			log.Printf("Failed `cel_ctl upload`: %v", err)
		}
	}()

	session, err := a.CreateSession(ctx, args, uc.UseBuiltins)
	if err != nil {
		return err
	}

	store := session.GetContext().GetObjectStore()

	var data []byte
	if data, err = ioutil.ReadFile(uc.File); err != nil {
		return err
	}

	objRef, err := store.PutNamedObject(filepath.Base(uc.File), data)
	if err != nil {
		return err
	}

	path := gcp.AbsoluteReference(session.GetBackend().GetBucket(), objRef)
	log.Printf("Succesfully uploaded file to storage: %v", path)

	if uc.Instance != "" {
		log.Printf("Downloading file from storage to '%v'.", uc.Instance)
		command := fmt.Sprintf("gsutil copy %s %s", path, uc.Destination)
		return RunCommandOnInstance(ctx, a, session, uc.Instance, command, defaultTimeout)
	}

	return nil
}
