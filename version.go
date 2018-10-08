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

package goomba

import (
	"bytes"
	"sync"
	"text/template"
)

const (
	// VersionTemplate is the tempate used to render the version information.
	VersionTemplate = `Goomba version {{.VersionSemVer}}{{if .VersionPreRelease}}-{{.VersionPreRelease}}{{end}}`

	// LongVersionTemplate is the tempate used to render the version
	// information.
	LongVersionTemplate = `Goomba version {{.VersionSemVer}}{{if .VersionPreRelease}}-{{.VersionPreRelease}}{{end}} build {{.VersionBuildID}} at {{.VersionTimestamp}}`
)

var (
	version *Version

	once sync.Once
)

// Version type defines the version information about the application.
type Version struct {
	// VersionSemVer is the sermver based number of the app
	VersionSemVer string

	// VersionBuildID is the latest commit hash
	VersionBuildID string

	// VersionTimestamp represents when the application was built
	VersionTimestamp string

	// VersionPreRelease is a pre-release tag for the application
	// like release-candidate, beta, dev, etc ...
	VersionPreRelease string
}

// ShowVersion shows the short version information.
func (v *Version) ShowVersion() (string, error) {
	buf := new(bytes.Buffer)

	t := template.Must(template.New("versionTemplate").Parse(VersionTemplate))
	err := t.Execute(buf, v)

	return buf.String(), err
}

// ShowLongVersion shows the long version information.
func (v *Version) ShowLongVersion() (string, error) {
	buf := new(bytes.Buffer)

	t := template.Must(template.New("versionTemplate").Parse(LongVersionTemplate))
	err := t.Execute(buf, v)

	return buf.String(), err
}

// GetVersion singleton pattern implementation
func GetVersion() *Version {
	once.Do(func() {
		version = &Version{}
	})

	return version
}
