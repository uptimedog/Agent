// Copyright 2020 Uptimedog. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
    "fmt"
    "time"

    log "github.com/sirupsen/logrus"
    "github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Start the uptimedog agent",
    Run: func(cmd *cobra.Command, args []string) {
        if Verbose {
            log.SetLevel(log.DebugLevel)
        }

        if APIKey == "" {
            panic("Uptimedog API Key is required")
        }

        if APIServer == "" {
            panic("Uptimedog API Server is required")
        }

        log.Info(fmt.Sprintf("Start the agent, Remote server is %s", APIServer))

        for {
            log.Info("Agent running")
            log.Debug("Add -v for verbose logs")
            time.Sleep(1 * time.Second)
        }
    },
}

func init() {
    rootCmd.AddCommand(runCmd)
}
