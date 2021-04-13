default: build

test:
	go test ./... -v

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-pingaccess ~/.tflint.d/plugins

checks:
	@go fmt ./...
	@staticcheck ./...
	@gosec ./...
	@goimports -w rules

