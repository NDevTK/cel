// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common_test

import (
	"chromium.googlesource.com/enterprise/cel/go/common"
	"fmt"
)

func ExampleTasks_Go_single() {
	j := common.NewTasks("my jobs")
	j.Go(func() error {
		fmt.Println("hello from job")
		return fmt.Errorf("error from job")
	})
	err := j.Join()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Output:
	// hello from job
	// error from job
}

func ExampleTasks_Go_parallel() {
	j := common.NewTasks(nil)
	for i := 0; i < 10; i++ {
		c := i
		j.Go(func() error {
			return fmt.Errorf("error from job %d", c+1)
		})
	}
	el := common.UnpackErrorList(j.Join())
	for _, e := range el {
		fmt.Println(e.Error())
	}

	// Unordered output:
	// error from job 1
	// error from job 2
	// error from job 3
	// error from job 4
	// error from job 5
	// error from job 6
	// error from job 7
	// error from job 8
	// error from job 9
	// error from job 10
}
