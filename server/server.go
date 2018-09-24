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
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/ansicolor"
	"github.com/goombaio/log"
)

var (
	loggerPrefixes = []string{
		ansicolor.ColorTrueColors("server", 39, 174, 96, 15, 15, 15),
		ansicolor.ColorTrueColors(time.Now().Format(time.RFC850), 41, 128, 185, 15, 15, 15),
	}
)

// Server ...
type Server struct {
	// Unique ID for the cluster node.
	// Used for traceability, metrics, monitoring, etc ...
	ID uuid.UUID

	// Name of the cluster node.
	Name string

	// logger is the custom log.Logger for the server
	logger log.Logger
}

// NewServer ...
func NewServer(name string) *Server {
	s := &Server{
		ID:     uuid.New(),
		Name:   name,
		logger: log.NewFmtLogger(os.Stderr),
	}

	return s
}

// Start ...
func (s *Server) Start() error {
	_ = s.logger.Log(loggerPrefixes, "Start", s.Name, "..")

	return nil
}

// Stop ...
func (s *Server) Stop() error {
	_ = s.logger.Log(loggerPrefixes, "Stop", s.Name, "..")

	return nil
}

// Run ...
func Run() error {
	server := NewServer("mainserver")

	go func() {
		handleSignals()
		_ = server.Stop()
	}()

	err := server.Start()

	return err
}

func handleSignals() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	fmt.Println(" ==> signal received")
}
