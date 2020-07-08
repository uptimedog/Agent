// Copyright 2020 Uptimedog. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Verbose var
var Verbose bool

// APIKey var
var APIKey string

var rootCmd = &cobra.Command{
	Use: "agent",
	Short: `Monitor internal services with uptimedog agent

Uptimedog and the agent is in early stages of development, and we'd love to hear your
feedback at <https://github.com/Uptimedog/agent>`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(
		&APIKey,
		"api_key",
		"k",
		"",
		"Your Uptimedog API Key",
	)
}

// Execute runs cmd tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
