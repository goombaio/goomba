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
	"net"
	"net/rpc"

	"github.com/goombaio/log"
)

// Server ...
type Server struct {
	config *Config
	logger log.Logger
}

// NewServer ...
func NewServer(logger log.Logger, config *Config) *Server {
	if !config.Enabled {
		return nil
	}

	server := &Server{
		logger: logger,
		config: config,
	}
	return server
}

// Start ...
func (s *Server) Start() error {
	s.logger.Info("Starting server..")

	listener, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		s.logger.Error(err, "can't start server")
		return err
	}

	msg := fmt.Sprintf("Server listening at %s.", s.config.Address)
	s.logger.Info(msg)

	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				s.logger.Error(err, "on incomming connection")
			}
			go rpc.ServeConn(client)
		}
	}()

	return nil
}

// Stop ...
func (s *Server) Stop() error {
	s.logger.Info("Stopping server..")

	return nil
}
