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

package server_test

import (
	"io/ioutil"
	"testing"

	"github.com/goombaio/goomba/server"
)

func TestNewServer(t *testing.T) {
	config := server.DefaultConfig()
	_ = server.NewServer(config)
}

func TestServer_Start(t *testing.T) {
	config := server.DefaultConfig()
	config.LogOutput = ioutil.Discard

	server := server.NewServer(config)

	go server.Start()
	server.Stop()
}

func TestServer_Restart(t *testing.T) {
	config := server.DefaultConfig()
	config.LogOutput = ioutil.Discard

	server := server.NewServer(config)

	go server.Start()
	go server.Restart()
	server.Stop()
}

func TestServer_Stop(t *testing.T) {
	config := server.DefaultConfig()
	config.LogOutput = ioutil.Discard

	server := server.NewServer(config)

	go server.Start()
	server.Stop()
}
