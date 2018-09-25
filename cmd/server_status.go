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
	"net/rpc"

	"github.com/goombaio/cli"
)

// ServerStatusCommand ...
var ServerStatusCommand *cli.Command

func init() {
	ServerStatusCommand = cli.NewCommand("status", "Get the status of the Goomba server")
	ServerStatusCommand.LongDescription = "status command get the status of the Goomba server node and cluster."
	ServerStatusCommand.Run = func(c *cli.Command) error {
		client, err := rpc.Dial("tcp", "localhost:7331")
		if err != nil {
			return err
		}
		defer client.Close()

		var reply string
		err = client.Call("StatusResponse.Status", "args", &reply)
		if err != nil {
			return err
		}

		fmt.Println("Result:", reply)

		return nil
	}
}
