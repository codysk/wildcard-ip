COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)
LDFLAGS := "\
-X github.com/codysk/wildcard-ip/cmd/version.GitCommit=$(COMMIT) \
-X github.com/codysk/wildcard-ip/cmd/version.GitTag=$(TAG) \
-linkmode 'external' \
-extldflags '-static' \
"
FLAGS := -ldflags $(LDFLAGS) -mod=vendor

all: build

build:
	CGO_ENABLED=1 go build $(FLAGS) -o bin/wildcard-ip github.com/codysk/wildcard-ip/cmd

clean:
	rm -rf bin
