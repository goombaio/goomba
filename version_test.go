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

package goomba_test

import (
	"fmt"
	"testing"

	"github.com/goombaio/goomba"
)

func TestVersion(t *testing.T) {
	version := &goomba.Version{}

	if version.SemVer != "" {
		t.Fatalf("SemVer expected to be blank but got %s", version.SemVer)
	}

	if version.BuildID != "" {
		t.Fatalf("BuildID expected to be blank but got %s", version.BuildID)
	}

	if version.Timestamp != "" {
		t.Fatalf("Timestamp expected to be blank but got %s", version.Timestamp)
	}

	if version.PreRelease != "" {
		t.Fatalf("Timestamp expected to be blank but got %s", version.PreRelease)
	}
}

func TestVersion_ShowVersion(t *testing.T) {
	semver := "1.2.3"
	buildid := "master-b61ad71"
	timestamp := "2018-09-17.06:58:24.UTC"
	prerelease := "dev"

	version := &goomba.Version{
		SemVer:     semver,
		BuildID:    buildid,
		Timestamp:  timestamp,
		PreRelease: prerelease,
	}

	// expectedVersionInfo := fmt.Sprintf("Goomba version %s-%s build %s at %s", semver, prerelease, buildid, timestamp)
	expectedVersionInfo := fmt.Sprintf("Goomba version %s-%s", semver, prerelease)
	versionInfo, err := version.ShowVersion()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if versionInfo != expectedVersionInfo {
		t.Fatalf("Expected %q but got %q", expectedVersionInfo, versionInfo)
	}
}

func TestVersion_ShowLongVersion(t *testing.T) {
	semver := "1.2.3"
	buildid := "master-b61ad71"
	timestamp := "2018-09-17.06:58:24.UTC"
	prerelease := "dev"

	version := &goomba.Version{
		SemVer:     semver,
		BuildID:    buildid,
		Timestamp:  timestamp,
		PreRelease: prerelease,
	}

	expectedVersionInfo := fmt.Sprintf("Goomba version %s-%s build %s at %s", semver, prerelease, buildid, timestamp)
	versionInfo, err := version.ShowLongVersion()
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	if versionInfo != expectedVersionInfo {
		t.Fatalf("Expected %q but got %q", expectedVersionInfo, versionInfo)
	}
}
