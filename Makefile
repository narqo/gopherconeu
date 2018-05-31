REPO := github.com/narqo/gopherconeu

GO   ?= go
GOOS ?= linux
GOARCH ?= amd64

APP := gophercon
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GITSHA := $(shell git rev-parse --short HEAD 2>/dev/null)

ifndef VERSION
	VERSION := git-$(GITSHA)
endif

PORT := 8000
HEALTH_PORT := 8001

CONTAINER_IMAGE := docker.io/varankinv/$(APP)

SHELL=/bin/bash

all: build

build:
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	$(GO) build \
		-o BUILD/$(APP) \
		-ldflags '\
			-X $(REPO)/version.COMMIT=$(GIT_SHA) \
			-X $(REPO)/version.BuildTime=$(BUILD_TIME) \
			-X $(REPO)/version.Version=$(VERSION)' \
		$(REPO)/cmd/$(APP)

container: build
	docker build -t $(CONTAINER_IMAGE):$(VERSION) .

test:
	$(GO) test ./...

run:
	PORT=$(PORT) HEATH_PORT=$(HEALTH_PORT) ./BUILD/$(APP)

clean:
	$(RM) -r BUILD/

