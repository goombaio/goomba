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

package service_test

import (
	"io/ioutil"
	"testing"

	"github.com/goombaio/goomba/service"
)

func TestNewNopService(t *testing.T) {
	config := service.DefaultConfig()
	_ = service.NewNopService(config)
}

func TestNewNopService_Start(t *testing.T) {
	config := service.DefaultConfig()
	config.LogOutput = ioutil.Discard
	service := service.NewNopService(config)

	err := service.Start()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewNopService_Restart(t *testing.T) {
	config := service.DefaultConfig()
	config.LogOutput = ioutil.Discard
	service := service.NewNopService(config)

	err := service.Restart()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewNopService_Stop(t *testing.T) {
	config := service.DefaultConfig()
	config.LogOutput = ioutil.Discard
	service := service.NewNopService(config)

	err := service.Stop()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNopServiceString(t *testing.T) {
	config := service.DefaultConfig()
	config.LogOutput = ioutil.Discard
	service := service.NewNopService(config)

	str := service.String()
	if str != "" {
		t.Fatalf("Expected blank string but got %s", str)
	}
}
