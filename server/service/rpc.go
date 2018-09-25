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

package service

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/google/uuid"
	"github.com/goombaio/log"
)

// RPCService type represents a service
type RPCService struct {
	// Service configuration
	config *Config

	// Unique service ID
	// Used for traceability, metrics, monitoring, etc ...
	ID uuid.UUID

	// Name of the service.
	Name string

	// logger is the custom log.Logger for the service.
	logger log.Logger
}

// NewRPCService creates a new server given a configuration.
func NewRPCService(config *Config) *RPCService {
	s := &RPCService{
		config: config,

		ID:   config.ID,
		Name: "rpc-server",

		logger: log.NewFmtLogger(config.LogOutput),
	}

	return s
}

// Start ...
func (rs *RPCService) Start() error {
	_ = rs.logger.Log(rs.config.LogPrefixes, "Start service", "-", rs.String())

	listener, err := net.Listen("tcp", "0.0.0.0:7331")
	if err != nil {
		_ = rs.logger.Log(rs.config.LogPrefixes, "ERROR:", err)
		return err
	}

	go rpc.Accept(listener)

	return nil
}

// Restart ...
func (rs *RPCService) Restart() error {
	_ = rs.logger.Log(rs.config.LogPrefixes, "Restart service", "-", rs.String())

	return nil
}

// Stop ...
func (rs *RPCService) Stop() error {
	_ = rs.logger.Log(rs.config.LogPrefixes, "Stop service", "-", rs.String())

	return nil
}

// String implements fmt.Stringer interface and returns the string
// representation of this type.
func (rs *RPCService) String() string {
	str := fmt.Sprintf("Name: %s - ID: %s", rs.Name, rs.ID)

	return str
}
