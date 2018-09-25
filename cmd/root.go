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
	"github.com/goombaio/cli"
)

// RootCommand is the main Command of the application.
var RootCommand = &cli.Command{
	Name:             "goomba",
	ShortDescription: "Goomba CLI",
	LongDescription:  "A workflow based data pipeline and ETL framework for golang. https://goomba.io",
	Run: func(c *cli.Command) error {
		c.Usage()

		return nil
	},
}

// Execute is the main entry point of the cli application.
//
// It will run  Command.Execute() base method that will parse for Commands,
// SubCommands, Flags, Arguments and route to the appropiate place if no error
// is found.
func Execute() error {
	err := RootCommand.Execute()
	if err != nil {
		_ = RootCommand.Logger().Log("ERROR:", err)
	}

	return nil
}
