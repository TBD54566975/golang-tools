// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringintconv_test

import (
	"testing"

	"github.com/TBD54566975/x/tools/go/analysis/analysistest"
	"github.com/TBD54566975/x/tools/go/analysis/passes/stringintconv"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, stringintconv.Analyzer, "a", "typeparams")
}
