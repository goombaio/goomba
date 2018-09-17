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
	"fmt"
	"os"

	"github.com/goombaio/cli"
	"github.com/goombaio/goomba"
	"github.com/goombaio/goomba/cmd"
	"github.com/goombaio/log"
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
	output := os.Stdout
	logger := log.NewFmtLogger(output)

	rootCommand := cli.NewCommand("goomba", "Goomba CLI")
	rootCommand.Run = func(c *cli.Command) error {
		c.Usage()

		return nil
	}

	versionCommand := cli.NewCommand("version", "Show version information")
	versionCommand.LongDescription = `version command shows the version 
  information about this program. It consists in 3 parts; the first one is a 
  canonical version following the semver specification. The second part is an 
  ID that identifies a single build which has the same versionn, currently we 
  use the git hash as this ID. And finally, the third part is a timestamp that 
  reflects when the project was built`
	versionCommand.Run = func(c *cli.Command) error {
		// TODO:
		// - Check if -h or --help is used and show subCommand Usage
		//   https://github.com/goombaio/goomba/issues/1
		// - By default only show SemVer & PreRelease is available
		// - Add --long support and show all info if it is used
		//   https://github.com/goombaio/goomba/issues/2
		//
		// c.Usage()

		version := &goomba.Version{
			SemVer:     VersionSemVer,
			BuildID:    VersionBuildID,
			Timestamp:  VersionTimestamp,
			PreRelease: VersionPreRelease,
		}
		versionInformation, err := version.ShowVersion()

		fmt.Println(versionInformation)

		return err
	}

	rootCommand.AddCommand(versionCommand)

	rootCommand.AddCommand(cmd.ServerCommand)

	err := rootCommand.Execute()
	if err != nil {
		logger.Log("ERROR:", err)
		os.Exit(1)
	}
}
