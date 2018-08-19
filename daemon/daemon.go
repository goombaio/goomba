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
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	signalChan chan os.Signal
	exitChan   chan int
}

// NewDaemon ...
func NewDaemon(logger log.Logger, config *Config) *Daemon {
	daemon := &Daemon{
		config:     config,
		logger:     logger,
		signalChan: make(chan os.Signal, 1),
		exitChan:   make(chan int, 1),
	}
	return daemon
}

// Start ...
func (d *Daemon) Start() {
	d.logger.Info("Start goomba..")

	d.setupServer()
	d.setupClient()

	d.handleSignals()
}

func (d *Daemon) setupServer() {
	// Check if server is enabled and run it
	if !d.config.Server.Enabled {
		d.logger.Info("Server is disabled.")
		return
	}

	d.logger.Info("Server enabled.")
	d.Server = server.NewServer(d.logger, d.config.Server)
	_ = d.Server.Start()
}

func (d *Daemon) setupClient() {
	// Check if client is enabled and run it
	if !d.config.Client.Enabled {
		d.logger.Info("Client is disabled.")
		return
	}

	d.logger.Info("Client enabled.")
	d.Client = client.NewClient(d.logger, d.config.Client)
}

// Stop ...
func (d *Daemon) Stop() {
	d.logger.Info("Stop goomba..")
	if d.config.Server.Enabled {
		_ = d.Server.Stop()
	}
	os.Exit(0)
}

// handleSignals blocks until we get an ex-causing signal.
//
// Listens for the following signals:
// - SIGINT: Triggers when the user types CRTL-c.
// 			 Kill -SIGINT XXX
// - SIGQUIT: Triggers when the user types CRTL-\.
// 			  Also core dumps.
// 			  Kill -SIGQUIT XXX
// - SIGTERM: generic signal used to cause program termination.
//			  The normal way to politely ask a program to terminate.
// 			  Kill -SIGTERM XXX
//
// See: http://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html#Termination-Signals
func (d *Daemon) handleSignals() {
	signal.Notify(d.signalChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for {
			sig := <-d.signalChan
			switch sig {
			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				msg := fmt.Sprintf("Trap signal SIGINT: %s.", sig)
				d.logger.Info(msg)
				d.Stop()

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				msg := fmt.Sprintf("Trap signal SIGQUIT: %s.", sig)
				d.logger.Info(msg)
				d.exitChan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				msg := fmt.Sprintf("Trace: Trap signal SIGTERM: %s.", sig)
				d.logger.Info(msg)
				d.exitChan <- 0

			default:
				msg := fmt.Sprintf("Unknown signal %s.", sig)
				d.logger.Info(msg)
				d.exitChan <- 1
			}
		}
	}()
	code := <-d.exitChan
	os.Exit(code)
}
