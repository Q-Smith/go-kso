# Make specific flags
MAKEFLAGS += --silent

# ------------------------------------------------------------------------------- #
# Global Variables
# ------------------------------------------------------------------------------- #

# Global Variables
GIT_SHA = $(shell git rev-parse --short HEAD)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)

APP_NAME = kso
APP_VERSION = 0.1.0

# Use linker flags to provide version/build settings
CGO = 0
GOARCH = amd64
LDFLAGS = -ldflags "-X=main.Version=${APP_VERSION} -X main.Branch=${GIT_BRANCH} -X main.SHA=${GIT_SHA}"

# ------------------------------------------------------------------------------- #
# Golang Commands
# ------------------------------------------------------------------------------- #

## Task: go-clean
.PHONY: go-clean
go-clean:
	echo "Running clean..."
	go clean -testcache -cache

## Task: go-build
go-build: go-build-linux go-build-windows

## Task: go-build-linux
.PHONY: go-build-linux
go-build-linux: go-clean
	echo "Compiling for Linux"
	GOOS=linux GOARCH=${GOARCH} CGO_ENABLED=${CGO} go build -v ${LDFLAGS} -o ./bin/kso-linux-${GOARCH} ./cmd/kso/

## Task: go-build-windows
.PHONY: go-build-windows
go-build-windows: go-clean
	echo "Compiling for Windows"
	GOOS=windows GOARCH=${GOARCH} CGO_ENABLED=${CGO} go build -v ${LDFLAGS} -o ./bin/kso-windows-${GOARCH}.exe ./cmd/kso/
