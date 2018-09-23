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

package client

import (
	"context"
	"net/rpc"

	"github.com/goombaio/goomba/server"
)

// RPCClient ...
type RPCClient struct {
	Client *rpc.Client
}

// Close terminates the underlying client.
func (rc *RPCClient) Close() error {
	if rc.Client != nil {
		err := rc.Client.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Status ...
func (rc *RPCClient) Status(ctx context.Context) (string, error) {
	var (
		request  = &server.Request{}
		response = new(server.Response)
	)

	err := rc.Client.Call("Status", request, response)
	if err != nil {
		return "", err
	}

	return response.Message, nil
}
