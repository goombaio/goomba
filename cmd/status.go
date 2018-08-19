// Copyright 2018, Goomba.io project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package goomba

import (
	"fmt"
	"os"

	"github.com/goombaio/goomba/daemon"

	"github.com/goombaio/log"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(StatusCmd)
}

// StatusCmd represents the start command
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Goomba status",
	Long:  `Show goomba goomba status`,
	Run: func(cmd *cobra.Command, args []string) {
		loggerOutput := os.Stderr
		_ = log.NewLogger(loggerOutput)

		config := daemon.DefaultConfig()
		// TODO: Merge default configuration with loaded configuration

		fmt.Println("Server:")
		fmt.Println("-------")
		fmt.Printf("Enabled: %v.\n", config.Server.Enabled)

		fmt.Println("Client:")
		fmt.Println("-------")
		fmt.Printf("Enabled: %v.\n", config.Client.Enabled)
	},
}
