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
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/ansicolor"
	"github.com/goombaio/log"
	"github.com/goombaio/namegenerator"
)

var (
	loggerPrefixes = []string{
		ansicolor.ColorTrueColors("server", 39, 174, 96, 15, 15, 15),
		ansicolor.ColorTrueColors(time.Now().Format(time.RFC850), 41, 128, 185, 15, 15, 15),
	}
)

// Server ...
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

// NewServer ...
func NewServer(config *Config) *Server {
	s := &Server{
		config: config,
		ID:     uuid.New(),
		Name:   "server-node",
		logger: log.NewFmtLogger(os.Stderr),
	}

	s.Name = s.generateRandomName()

	return s
}

// Start ...
func (s *Server) Start() error {
	_ = s.logger.Log(loggerPrefixes, "Start Goomba server -", "Name:", s.Name, "ID:", s.ID, "..")

	// Listen for syscall signals to gracefully stop the server
	s.handleSignals()

	// Block forever
	select {}
}

// Reload ...
func (s *Server) Reload() error {
	_ = s.logger.Log(loggerPrefixes, "Reload Goomba server -", "Name:", s.Name, "ID:", s.ID, "..")

	return nil
}

// Restart ...
func (s *Server) Restart() error {
	_ = s.logger.Log(loggerPrefixes, "Restart Goomba server -", "Name:", s.Name, "ID:", s.ID, "..")

	return nil
}

// Stop ...
func (s *Server) Stop() error {
	_ = s.logger.Log(loggerPrefixes, "Stop Goomba server -", "Name:", s.Name, "ID:", s.ID, "..")

	return nil
}

// handleSignal ...
func (s *Server) handleSignals() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			sig := <-signalChan
			switch sig {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				_ = s.logger.Log(loggerPrefixes, "hungup")
				_ = s.Reload()
				_ = s.Restart()

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				_ = s.logger.Log(loggerPrefixes, "interrupt")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				_ = s.logger.Log(loggerPrefixes, "force stop")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				_ = s.logger.Log(loggerPrefixes, "stop and core dump")
				_ = s.Stop()
				exitChan <- 0

			default:
				_ = s.logger.Log(loggerPrefixes, "Unknown signal.")
				_ = s.Stop()
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

// generateRandomName ...
func (s *Server) generateRandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	generatedName := nameGenerator.Generate()

	return generatedName
}
