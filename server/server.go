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
	"net"
	"net/rpc"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/goombaio/log"
)

var (
	loggerPrefixes = []string{
		color.BlueString("server"),
		color.GreenString(time.Now().Format(time.RFC850)),
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

	listener net.Listener
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
	s.logger.Log(loggerPrefixes, "Start", s.Name, "..")

	rpc.RegisterName("Server", &RPCServer{})

	listener, err := net.Listen("tcp", "0.0.0.0:7331")
	if err != nil {
		return err
	}
	s.listener = listener

	rpc.Accept(s.listener)

	return nil
}

// Stop ...
func (s *Server) Stop() error {
	s.logger.Log(loggerPrefixes, "Stop", s.Name, "..")

	if s.listener != nil {
		err := s.listener.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
