// Copyright 2020 Uptimedog. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/uptimedog/agent/pkg"
)

// TestPkgName test cases
func TestPkgName(t *testing.T) {
	pkg.Expect(t, "Agent", "Agent")
}
