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
	"os"

	"github.com/goombaio/goomba/daemon"

	"github.com/goombaio/log"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(StartCmd)
}

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start goomba",
	Long:  `Start goomba services reading configuration`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		loggerOutput := os.Stderr
		logger := log.NewLogger(loggerOutput)

		config := daemon.DefaultConfig()

		daemon := daemon.NewDaemon(logger, config)

		err := daemon.Run()
		if err != nil {
			logger.Error(err, "running goomba")
			os.Exit(1)
		}
	},
}
