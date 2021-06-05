// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"testing"

	"github.com/uptimedog/agent/pkg"

	"github.com/franela/goblin"
)

// TestUnitOsRelease
func TestUnitOsRelease(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("#OsRelease", func() {
		g.It("It should satisfy test cases", func() {
			path := fmt.Sprintf(
				"%s/cache/os-release",
				pkg.GetBaseDir("cache"),
			)

			result, err := GetOsRelease([]string{path})

			g.Assert(result.Name).Equal("Ubuntu")
			g.Assert(err).Equal(nil)

			result, err = GetOsRelease([]string{})

			g.Assert(result.Name).Equal("")
			g.Assert(err.Error()).Equal("open /etc/os-release: no such file or directory")
		})
	})
}
