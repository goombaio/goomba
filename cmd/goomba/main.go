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

package main

import (
	"fmt"

	goomba "github.com/goombaio/goomba/cmd"
)

var (
	// Version is the current version number
	Version string
	// Build is the current build id
	Build string
)

func main() {
	goomba.RootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}`)
	goomba.RootCmd.Version = showVersionInfo(Version, Build)

	goomba.Execute()
}

// showVersionInfo returns version and build information
func showVersionInfo(version, build string) string {
	tpl := "version %s build %s\n"
	output := fmt.Sprintf(tpl, version, build)
	return output
}
