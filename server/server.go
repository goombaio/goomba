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

	"github.com/google/uuid"
	"github.com/goombaio/goomba/service"
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

	// services are the services this server managees
	services []service.Service

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

		services: make([]service.Service, 0),
	}

	s.logger.Log(s.config.LogPrefixes, "New server", "-", s.String())

	return s
}

// Start starts a server and  its belonging services.
func (s *Server) Start() error {
	s.logger.Log(s.config.LogPrefixes, "Start server", "-", s.String())

	// Start services
	for _, registeredService := range s.services {
		go func(registeredService service.Service) {
			_ = registeredService.Start()
		}(registeredService)
	}

	// Listen for syscall signals to gracefully stop the server
	s.handleSignals()

	return nil
}

// Restart reloads and restarts a server and its belonging services.
func (s *Server) Restart() error {
	// Restart services
	for _, service := range s.services {
		err := service.Restart()
		if err != nil {
			return err
		}
	}

	s.logger.Log(s.config.LogPrefixes, "Restart server", "-", s.String())

	return nil
}

// Stop stops a server belonging services and the server itself.
func (s *Server) Stop() error {
	// Stop services
	for _, service := range s.services {
		err := service.Stop()
		if err != nil {
			return err
		}
	}

	s.logger.Log(s.config.LogPrefixes, "Stop server", "-", s.String())

	return nil
}

// Services returns the service list managed by this server
func (s *Server) Services() []service.Service {
	return s.services
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
				s.logger.Log(s.config.LogPrefixes, "hungup")
				_ = s.Restart()

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				s.logger.Log(s.config.LogPrefixes, "interrupt")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				s.logger.Log(s.config.LogPrefixes, "force stop")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				s.logger.Log(s.config.LogPrefixes, "stop and core dump")
				_ = s.Stop()
				exitChan <- 0

			default:
				s.logger.Log(s.config.LogPrefixes, "Unknown signal.")
				_ = s.Stop()
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

// RegisterService registers a service in this server, this server then will
// will control servicep's lifecycle.
func (s *Server) RegisterService(service service.Service) error {
	s.services = append(s.services, service)

	s.logger.Log(s.config.LogPrefixes, "Register service", "-", service.String())

	return nil
}

// String implements fmt.Stringer interface and returns the string
// representation of this type.
func (s *Server) String() string {
	str := fmt.Sprintf("Name: %s - ID: %s", s.Name, s.ID)

	return str
}
