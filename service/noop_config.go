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
	"io"
	"io/ioutil"

	"github.com/goombaio/guid"
)

// NoopConfig type represents a server configuration.
type NoopConfig struct {
	ID          string
	Name        string
	LogOutput   io.Writer
	LogPrefixes []string
}

// DefaultConfig returns the server default configuration.
func DefaultConfig() *NoopConfig {
	c := &NoopConfig{
		ID:          guid.New(),
		Name:        "Noop-service",
		LogOutput:   ioutil.Discard,
		LogPrefixes: []string{},
	}

	return c
}
