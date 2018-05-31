REPO := github.com/narqo/gopherconeu

GO   ?= go
#GOOS ?= linux
#GOARCH ?= amd64

APP := gophercon

all: build

build:
	$(GO) build -o BUILD/$(APP) $(REPO)/cmd/$(APP)

test:
	$(GO) test ./...

run:
	PORT=$(PORT) ./BUILD/$(APP)

clean:
	$(RM) -r BUILD/

