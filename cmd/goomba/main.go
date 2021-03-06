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

	"github.com/goombaio/cli"
	"github.com/goombaio/goomba"
	"github.com/goombaio/goomba/cmd"
)

var (
	// VersionSemVer ...
	VersionSemVer string

	// VersionBuildID ...
	VersionBuildID string

	// VersionTimestamp ...
	VersionTimestamp string

	// VersionPreRelease ...
	VersionPreRelease string
)

func main() {
	v := goomba.GetVersion()
	v.VersionSemVer = VersionSemVer
	v.VersionBuildID = VersionBuildID
	v.VersionTimestamp = VersionTimestamp
	v.VersionPreRelease = VersionPreRelease

	/* Setup commands, sub-commands, flags and add them to the RootCommand. */

	// server
	cmd.ServerCommand.AddCommand(cmd.ServerStartCommand)
	cmd.ServerCommand.AddCommand(cmd.ServerStatusCommand)
	cmd.RootCommand.AddCommand(cmd.ServerCommand)

	// version
	cmd.RootCommand.AddCommand(cmd.VersionCommand)

	// Run cli.Execute() method and force exit if error.
	err := cli.Execute(cmd.RootCommand)
	if err != nil {
		cmd.RootCommand.Logger().Log("ERROR:", err)
		os.Exit(1)
	}
}
