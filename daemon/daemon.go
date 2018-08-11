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

package daemon

import (
	"github.com/goombaio/goomba/client"
	"github.com/goombaio/goomba/server"
	"github.com/goombaio/log"
)

// Daemon ...
type Daemon struct {
	config *Config
	logger log.Logger

	Client *client.Client
	Server *server.Server
}

// NewDaemon ...
func NewDaemon(logger log.Logger, config *Config) *Daemon {
	daemon := &Daemon{
		config: config,
		logger: logger,
	}

	logger.Info("Starting goomba..")
	return daemon
}

// Run ...
func (d *Daemon) Run() error {
	d.setupServer()
	d.setupClient()

	return nil
}

func (d *Daemon) setupServer() {
	// Check if server is enabled and run it
	if !d.config.Server.Enabled {
		d.logger.Info("Server is disabled.")
		return
	}

	d.logger.Info("Server enabled..")
	d.Server = server.NewServer(d.logger, d.config.Server)
}

func (d *Daemon) setupClient() {
	// Check if client is enabled and run it
	if !d.config.Client.Enabled {
		d.logger.Info("Client is disabled.")
		return
	}

	d.logger.Info("Client enabled..")
	d.Client = client.NewClient(d.logger, d.config.Client)
}
