// +build tools

// This file induces Go modules to add mjibson/esc as an external dependency.
// (https://stackoverflow.com/questions/52428230/how-do-go-modules-work-with-installable-commands)
// mjibson/esc is a file embedder that build.py uses to compile CELab binaries.
// It must be in the ./vendor/github.com directory before compilation starts.

package tools

import (
	_ "github.com/mjibson/esc"
)
