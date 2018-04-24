// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

// List of known annotator generators. We ignore the package, and match all
// functions that match the name.
//
// If you are here to update the list of known annotators, then this is the
// list you are looking for.
var annotators = map[string]bool{
	"LoggedAction":           true,
	"GcpLoggedAction":        true,
	"GcpLoggedServiceAction": true,
}

func usage(f *flag.FlagSet) {
	fmt.Fprint(os.Stderr, `vet_annotations: Vet known annotation invocations.

Function calls like common.LoggedAction() are supposed to be made using a
special form as follows:

     func DoSomething(l Logger, i int) (err error) {
       defer LoggedAction(l, &err, "create some sort of object")()
       //                                                       ^^
       //          Note the additional function invocation here !!
       //          -----------------------------------------------
       DoStuff()

This tool ensures that invocations of such annotators are correct. The list of
annotators that are recognized by this tool are:
`)
	for a, _ := range annotators {
		fmt.Fprintf(os.Stderr, "    %s\n", a)
	}
	fmt.Fprint(os.Stderr, `
The list of files is defined in go/tools/vet_annotations/vet.go . Please update
that file if any are missing.
`)
	os.Exit(1)
}

type message struct {
	pos     token.Position
	message string
}

// isDeferIgnoringFuncReturn returns true if the given ast.Node represents a
// 'defer' statement that ignores the result of an annotator function
// generator.
//
// We are looking for things like: defer targetFunction(...) where
// targetFunction is one of the functions we don't want to ignore return values
// from.
//
// So our AST starting from the defer should be something like:
//
//     DeferStmt
//      |
//      +--Call -> CallExpr                                .... Outer call
//                  |
//                  +-- Fun -> Expr is-a CallExpr          .... Inner call
//                              |
//                              +-- Fun -> Expr is a reference to an annotator generator.
//
// The anomalous case is one where we have an AST as the above, but there is no
// outer call. I.e.:
//
//    defer targetFunction(...)
//
// Instead of
//
//    defer targetFunction(...)()
//
func isDeferIgnoringFuncReturn(fset *token.FileSet, n ast.Node) (msgs []message) {
	d, ok := n.(*ast.DeferStmt)
	if !ok {
		return
	}

	// Not the Anomalous case
	if !isCallToAnnotator(d.Call) {
		return
	}

	return []message{{
		pos:     fset.Position(d.Pos()),
		message: fmt.Sprintf("defer statement discarding returned function")}}
}

// isCallToAnnotator returns true if the given CallExpr invokes an annotator.
func isCallToAnnotator(c *ast.CallExpr) (found bool) {
	return isRefToAnnotator(c.Fun)
}

// isRefToAnnotator returns true if the given ast.Node is a reference to an
// annotator function.
func isRefToAnnotator(n ast.Node) bool {
	switch c := n.(type) {
	case *ast.Ident:
		_, ok := annotators[c.Name]
		return ok

	case *ast.SelectorExpr:
		return isRefToAnnotator(c.Sel)
	}
	return false
}

// inspectPackage inspects a parsed Package for known issues.
func inspectPackage(fset *token.FileSet, name string, p *ast.Package) (msgs []message) {
	ast.Inspect(p, func(n ast.Node) bool {
		m := isDeferIgnoringFuncReturn(fset, n)
		if len(m) != 0 {
			msgs = append(msgs, m...)
		}
		return true
	})
	return
}

// inspectDir parses all the go code in a specified directory and inspects the
// generated ASTs for known issues. Each issue that's found is returend as a
// message element.
func inspectDir(dir string) (msgs []message, err error) {
	var fset token.FileSet
	pkgs, err := parser.ParseDir(&fset, dir, nil, parser.AllErrors)
	if err != nil {
		return
	}

	for n, p := range pkgs {
		m := inspectPackage(&fset, n, p)
		if len(m) != 0 {
			msgs = append(msgs, m...)
		}
	}

	return msgs, nil
}

func fatalf(err error, action string, v ...interface{}) {
	p_name := filepath.Base(os.Args[0])
	if err == nil {
		log.Printf("%s (%s)", p_name, fmt.Sprintf(action, v...))
	} else {
		log.Printf("%s (%s): %s", p_name, fmt.Sprintf(action, v...), err.Error())
	}
	os.Exit(1)
}

func main() {
	var dirs []string
	var showHelp bool

	flags := flag.NewFlagSet("vet_annotations", flag.ExitOnError)
	flags.SetOutput(os.Stderr)
	flags.BoolVar(&showHelp, "h", false, "show help")

	err := flags.Parse(os.Args[1:])

	if showHelp || err == flag.ErrHelp || len(flags.Args()) == 0 {
		usage(flags)
	}

	var msgs []message
	dirs = flags.Args()
	for _, dir := range dirs {
		m, err := inspectDir(dir)
		if err != nil {
			fatalf(err, "inspecting %s", dir)
		}
		if len(m) != 0 {
			msgs = append(msgs, m...)
		}
	}

	for _, m := range msgs {
		fmt.Printf("%s: %s\n", m.pos, m.message)
	}
}
