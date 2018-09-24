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
	"github.com/goombaio/goomba/server"
)

// ServerStartCommand ...
var ServerStartCommand *cli.Command

func init() {
	ServerStartCommand = cli.NewCommand("start", "Start a Goomba server")
	ServerStartCommand.LongDescription = `start command starts a Goomba server 
  node and runs until an interrupt is received. The server represents a single 
  node in a cluster.`
	ServerStartCommand.Run = func(c *cli.Command) error {
		err := server.Run()

		return err
	}
}
