// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This file provides an example command for static checkers
// conforming to the github.com/TBD54566975/golang-tools/go/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//	$ go build -o unitchecker main.go
//	$ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"github.com/TBD54566975/golang-tools/go/analysis/unitchecker"

	"github.com/TBD54566975/golang-tools/go/analysis/passes/appends"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/asmdecl"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/assign"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/atomic"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/bools"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/buildtag"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/cgocall"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/composite"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/copylock"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/directive"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/errorsas"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/framepointer"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/httpresponse"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/ifaceassert"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/loopclosure"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/lostcancel"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/nilfunc"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/printf"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/shift"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/sigchanyzer"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/stdmethods"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/stringintconv"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/structtag"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/testinggoroutine"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/tests"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/timeformat"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/unmarshal"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/unreachable"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/unsafeptr"
	"github.com/TBD54566975/golang-tools/go/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		appends.Analyzer,
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		testinggoroutine.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
