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
	"text/template"
)

const (
	// VersionTemplate is the tempate used to render the version information.
	VersionTemplate = `Goomba version {{.SemVer}}{{if .PreRelease}}-{{.PreRelease}}{{end}}`

	// LongVersionTemplate is the tempate used to render the version
	// information.
	LongVersionTemplate = `Goomba version {{.SemVer}}{{if .PreRelease}}-{{.PreRelease}}{{end}} build {{.BuildID}} at {{.Timestamp}}`
)

// Version ...
type Version struct {
	// SemVer is the current version number following the SemVer
	// specification.
	SemVer string

	// BuildID is the current build ID (usually the latest git commit hash).
	BuildID string

	// Timestamp is the timestamp when this application have been build.
	Timestamp string

	// PreRelease is the pre-release tag string of this application if it was
	// provided.
	PreRelease string
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
