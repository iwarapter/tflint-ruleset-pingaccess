TAG?=main
PKG=github.com/iwarapter/tflint-ruleset-pingaccess
REPO_INFO=$(shell git config --get remote.origin.url)

ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

LDFLAGS=-X main.version=$(TAG) -X $(PKG)/rules.version=$(TAG)

default: build

test:
	go test ./... -v

build:
	go build -v -ldflags="${LDFLAGS}"

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-pingaccess ~/.tflint.d/plugins

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w rules

