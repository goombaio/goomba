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

	"github.com/goombaio/ansicolor"
	"github.com/goombaio/goomba/service"
	"github.com/goombaio/log"
)

// Server type represents the main server responsible to start, stop or restart
// services as needed.
type Server struct {
	// Server configuration
	config *Config

	// Unique GUID for the cluster node.
	// Used for traceability, metrics, monitoring, etc ...
	GUID string

	// Name of the cluster node.
	Name string

	// services are the services this server managees
	services []service.Service

	// logger is the custom log.Logger for the server
	logger *log.ContextLogger
}

// NewServer creates a new server given a configuration.
func NewServer(config *Config) *Server {
	s := &Server{
		config: config,

		GUID: config.GUID,
		Name: config.Name,

		logger: log.NewContextLogger(config.LogOutput),

		services: make([]service.Service, 0),
	}

	s.logger.AddPrefix(ansicolor.ColorTrueColors(fmt.Sprintf("%T", s), 39, 174, 96, 15, 15, 15))
	s.logger.AddPrefix(ansicolor.ColorTrueColors(time.Now().Format(time.RFC850), 41, 128, 185, 15, 15, 15))

	s.logger.Log("New server", "-", s.String())

	return s
}

// Start starts a server and  its belonging services.
func (s *Server) Start() error {
	s.logger.Log("Start server", "-", s.String())

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

	s.logger.Log("Restart server", "-", s.String())

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

	s.logger.Log("Stop server", "-", s.String())

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
				s.logger.Log("hungup")
				_ = s.Restart()

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				s.logger.Log("interrupt")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				s.logger.Log("force stop")
				_ = s.Stop()
				exitChan <- 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				s.logger.Log("stop and core dump")
				_ = s.Stop()
				exitChan <- 0

			default:
				s.logger.Log("Unknown signal.")
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

	s.logger.Log("Register service", "-", service.String())

	return nil
}

// String implements fmt.Stringer interface and returns the string
// representation of this type.
func (s *Server) String() string {
	str := fmt.Sprintf("Name: %s - GUID: %s", s.Name, s.GUID)

	return str
}
