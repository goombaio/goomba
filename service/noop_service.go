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
	"github.com/goombaio/log"
)

// NoopService type represents a service that does nothing
type NoopService struct {
	// Service configuration
	config *NoopConfig

	// Unique service ID
	// Used for traceability, metrics, monitoring, etc ...
	ID string

	// Name of the service.
	Name string

	// logger is the custom log.Logger for the service.
	logger log.Logger
}

// NewNoopService creates a new service given a configuration.
func NewNoopService(config *NoopConfig) *NoopService {
	s := &NoopService{
		config: config,

		ID:   config.ID,
		Name: "Noop-service",

		logger: log.NewFmtLogger(config.LogOutput),
	}

	return s
}

// Start ...
func (rs *NoopService) Start() error {
	return nil
}

// Restart ...
func (rs *NoopService) Restart() error {
	return nil
}

// Stop ...
func (rs *NoopService) Stop() error {
	return nil
}

// String implements fmt.Stringer interface and returns the string
// representation of this type.
func (rs *NoopService) String() string {
	str := ""

	return str
}
