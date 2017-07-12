JOBDATE		?= $(shell date -u +%Y-%m-%dT%H%M%SZ)
GIT_REVISION	= $(shell git rev-parse --short HEAD)
VERSION		?= $(GIT_REVISION)

LDFLAGS		+= -X github.com/croomes/fiord/version.Version=$(VERSION)
LDFLAGS		+= -X github.com/croomes/fiord/version.Revision=$(GIT_REVISION)
LDFLAGS		+= -X github.com/croomes/fiord/version.BuildDate=$(JOBDATE)

.PHONY: release

test:
	go test -v `go list ./... | egrep -v /vendor/`

build:
	@echo "++ Building fiord"
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags "$(LDFLAGS)" -o fiord .