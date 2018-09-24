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

	"github.com/google/uuid"
	"github.com/goombaio/log"
)

// NopService type represents a service that does nothing
type NopService struct {
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

// NewNopService creates a new service given a configuration.
func NewNopService(config *Config) *NopService {
	s := &NopService{
		config: config,

		ID:   config.ID,
		Name: "nop-service",

		logger: log.NewFmtLogger(config.LogOutput),
	}

	return s
}

// Start ...
func (rs *NopService) Start() error {
	return nil
}

// Restart ...
func (rs *NopService) Restart() error {
	return nil
}

// Stop ...
func (rs *NopService) Stop() error {
	return nil
}

// String implements fmt.Stringer interface and returns the string
// representation of this type.
func (rs *NopService) String() string {
	str := fmt.Sprintf("Name: %s - ID: %s", rs.Name, rs.ID)

	return str
}
