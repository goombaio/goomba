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

package main

import (
	"os"

	"github.com/goombaio/goomba/cmd"
	"github.com/goombaio/log"
)

var (
	// VersionSemVer is the current version number following the SemVer
	// specification.
	VersionSemVer string

	// VersionBuildID is the current build ID (usually the latest git commit
	// hash).
	VersionBuildID string

	// VersionTimestamp is the timestamp when this application have been build.
	VersionTimestamp string

	// VersionPreRelease is the pre-release tag string of this application if
	// it was provided.
	VersionPreRelease string
)

func main() {
	// Setup main logger.
	output := os.Stdout
	logger := log.NewFmtLogger(output)

	/* Setup commands, subcommands and add them to the RootCommand. */

	// server
	cmd.ServerCommand.AddCommand(cmd.ServerStartCommand)
	cmd.ServerCommand.AddCommand(cmd.ServerStatusCommand)
	cmd.RootCommand.AddCommand(cmd.ServerCommand)

	// version
	cmd.RootCommand.AddCommand(cmd.VersionCommand)

	// Execute the RootCommand and force exit if error.
	err := cmd.Execute()
	if err != nil {
		logger.Log("ERROR:", err)
		os.Exit(1)
	}
}
