BINARY=goomba
MAIN_PACKAGE=cmd/${BINARY}/main.go
PACKAGES = $(shell go list ./...)
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
TIMESTAMP=`date -u '+%Y-%m-%d.%H:%M:%S.%Z'`
DIST_BUILD=bin
DIST_FOLDER=dist
DIST_INCLUDE_FILES=README.md LICENSE VERSION AUTHORS CONTRIBUTORS

# Setup -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.VersionSemVer=${VERSION} -X main.VersionBuildID=${BUILD} -X main.VersionTimestamp=${TIMESTAMP}"

# Build & Install

install: clean
install:		## Build and install package on your system (under $GOBIN)
	go install $(LDFLAGS) -v $(PACKAGES)

build: clean
build:			## Build binary
	mkdir -p bin
	go build $(LDFLAGS) -o ${DIST_BUILD}/${BINARY} ${MAIN_PACKAGE}

.PHONY: version
version:		## Show version information
	@echo "Goomba version $(VERSION) build $(BUILD) at ${TIMESTAMP}"

# Testing

.PHONY: test
test:			## Execute package tests
	go test -v $(PACKAGES)

.PHONY: test-race
test-race:
	go test -race -v $(PACKAGES)

cover-profile:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out

cover: cover-profile
cover: 			## Generate test coverage data
	go tool cover -func=coverage-all.out

cover-html: cover-profile
cover-html:		## Generate coverage report
	go tool cover -html=coverage-all.out

.PHONY: codecov
codecov:
	bash <(curl -s https://codecov.io/bash)

# BenchMarking

.PHONY: benchmark
benchmark:		## Execute package benchmarks
	go test -v $(PACKAGES) -benchmem -bench .

# Dependencies

deps:			## Install build dependencies
	go get -u
	go mod tidy -v
	go mod download
	go mod verify

dev-deps: deps
dev-deps:		## Install dev and build dependencies

clean:			## Delete generated development environment
	go clean
	rm -rf bin
	rm -rf coverage-all.out

# Lint

.PHONY: lint
lint:			## Lint source code
	./lint.bash

# Docs

godoc-serve:		## Serve documentation (godoc format) for this package at port HTTP 9090
	godoc -http=":9090"

# Distribution

dist-clean: clean
dist-clean:		## Delete generated development environment
	rm -rf dist

dist: dist-linux dist-darwin dist-windows dist-freebsd dist-openbsd
dist:			## Generate distribution binaries and packages

dist-linux-386:
	$(eval GOOS=linux)
	$(eval GOARCH=386)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-linux-amd64:
	$(eval GOOS=linux)
	$(eval GOARCH=amd64)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-linux-arm:
	$(eval GOOS=linux)
	$(eval GOARCH=arm)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-linux: dist-linux-386 dist-linux-amd64 dist-linux-arm

dist-darwin-386:
	$(eval GOOS=darwin)
	$(eval GOARCH=386)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-darwin-amd64:
	$(eval GOOS=darwin)
	$(eval GOARCH=amd64)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-darwin: dist-darwin-386 dist-darwin-amd64

dist-windows-386:
	$(eval GOOS=windows)
	$(eval GOARCH=386)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-windows-amd64:
	$(eval GOOS=windows)
	$(eval GOARCH=amd64)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.exe ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-windows: dist-windows-386 dist-windows-amd64

dist-freebsd-386:
	$(eval GOOS=freebsd)
	$(eval GOARCH=386)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-freebsd-amd64:
	$(eval GOOS=freebsd)
	$(eval GOARCH=amd64)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-freebsd-arm:
	$(eval GOOS=freebsd)
	$(eval GOARCH=arm)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-freebsd: dist-freebsd-386 dist-freebsd-amd64 dist-freebsd-arm

dist-openbsd-386:
	$(eval GOOS=openbsd)
	$(eval GOARCH=386)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-openbsd-amd64:
	$(eval GOOS=openbsd)
	$(eval GOARCH=amd64)
	mkdir -p ${DIST_FOLDER}/${GOOS}-${GOARCH}/
	go build -v ${LDFLAGS} -o ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${MAIN_PACKAGE}
	chmod +x ${BINARY}-${VERSION}-${GOOS}-${GOARCH}
	zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH}.zip ${BINARY}-${VERSION}-${GOOS}-${GOARCH} ${DIST_INCLUDE_FILES}
	mv ${BINARY}-${VERSION}* ${DIST_FOLDER}/${GOOS}-${GOARCH}/

dist-openbsd: dist-openbsd-386 dist-openbsd-amd64

include Makefile.help.mk
