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

package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/goombaio/log"
)

// Server type represents the main server responsible to start, stop or restart
// services as needed.
type Server struct {
	// Server configuration
	config *Config

	// Unique ID for the cluster node.
	// Used for traceability, metrics, monitoring, etc ...
	ID uuid.UUID

	// Name of the cluster node.
	Name string

	// logger is the custom log.Logger for the server
	logger log.Logger
}

// NewServer creates a new server given a configuration.
func NewServer(config *Config) *Server {
	s := &Server{
		config: config,

		ID:   config.ID,
		Name: config.Name,

		logger: log.NewFmtLogger(config.LogOutput),
	}

	return s
}

// Start starts a server and  its belonging services.
func (s *Server) Start() error {
	_ = s.logger.Log(s.config.LogPrefixes, "Start server", "-", "Name:", s.Name, "-", "ID:", s.ID)

	// Listen for syscall signals to gracefully stop the server
	s.handleSignals()

	// Start services

	return nil
}

// Restart reloads and restarts a server and its belonging services.
func (s *Server) Restart() error {
	_ = s.logger.Log(s.config.LogPrefixes, "Restart server", "-", "Name:", s.Name, "-", "ID:", s.ID)

	// Reload services

	// Restart services

	return nil
}

// Stop stops a server belonging services and the server itself.
func (s *Server) Stop() error {
	_ = s.logger.Log(s.config.LogPrefixes, "Stop server", "-", "Name:", s.Name, "-", "ID:", s.ID)

	// Stop services

	return nil
}

// handleSignal listens for syscall signals to gracefully stop, or restart a
// server and itss belonging services.
func (s *Server) handleSignals() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			sig := <-signalChan
			switch sig {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				_ = s.logger.Log(s.config.LogPrefixes, "hungup")
				_ = s.Restart()

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				_ = s.logger.Log(s.config.LogPrefixes, "interrupt")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				_ = s.logger.Log(s.config.LogPrefixes, "force stop")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				_ = s.logger.Log(s.config.LogPrefixes, "stop and core dump")
				_ = s.Stop()
				exitChan <- 0

			default:
				_ = s.logger.Log(s.config.LogPrefixes, "Unknown signal.")
				_ = s.Stop()
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
