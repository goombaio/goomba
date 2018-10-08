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

package cmd

import (
	"fmt"

	"github.com/goombaio/cli"
	"github.com/goombaio/goomba"
)

// VersionCommand is a cli.Command that shows version information about the
// program.
var VersionCommand = &cli.Command{
	Name:             "version",
	ShortDescription: "Show version information",
	LongDescription: `version command shows the version 
	information about this program. It consists in 3 parts; the first one is a 
	canonical version following the semver specification. The second part is an 
	ID that identifies a single build which has the same versionn, currently we 
	use the git hash as this ID. And finally, the third part is a timestamp that 
	reflects when the project was built.`,
	Run: func(c *cli.Command) error {
		version := goomba.GetVersion()

		longFlag := c.FlagName("-long")
		if longFlag.Parsed {
			result, err := version.ShowLongVersion()
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(c.Output(), "%s\n", result)

			return err
		}

		result, err := version.ShowVersion()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(c.Output(), "%s\n", result)

		return err
	},
}

func init() {
	LongVersionFlag := &cli.Flag{
		ShortName:   "-l",
		LongName:    "-long",
		Description: "Show long version information",
		Value:       "false",
		Parsed:      false,
	}

	VersionCommand.AddFlag(LongVersionFlag)
}
