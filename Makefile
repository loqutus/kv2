SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME=kv2

build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/server/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-client-linux-amd64 cmd/client/main.go
	CGO_ENABLED=0 GOARM=6 GOARCH=arm GOOS=linux go build -o bin/${BINARY_NAME}-linux-arm cmd/server/main.go
	CGO_ENABLED=0 GOARM=6 GOARCH=arm GOOS=linux go build -o bin/${BINARY_NAME}-client-linux-arm cmd/client/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-arm64 cmd/server/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o bin/${BINARY_NAME}-client-linux-arm64 cmd/client/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-mac-arm64 cmd/server/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-client-mac-arm64 cmd/client/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-mac-amd64 cmd/server/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-client-mac-amd64 cmd/client/main.go
	chmod +x bin/*

build_mac_arm64:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-mac-arm64 cmd/server/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-client-mac-arm64 cmd/client/main.go
	chmod +x bin/*

test_mac_arm64:
	./bin/kv2-mac-arm64 &
	KV2_LISTEN_PORT_CLIENT=6971 KV2_LISTEN_PORT_SERVER=6972 ./bin/kv2-mac-arm64 &
	sleep 1
	cd pkg/client/client && go test -count=1 -v -bench ./ || true
	for i in $$(ps aux | grep kv2 | grep -v grep | awk '{print $$2}'); do kill $$i; done

run:
	./bin/kv2-linux-arm64 &
	KV2_LISTEN_PORT_CLIENT=6971 KV2_LISTEN_PORT_SERVER=6972 ./bin/kv2-linux-arm64 &

run_mac_arm64:
	./bin/kv2-mac-arm64 &
	KV2_LISTEN_PORT_CLIENT=6971 KV2_LISTEN_PORT_SERVER=6972 ./bin/kv2-mac-arm64 &

test:
	cd pkg/client/client && go test -count=1 -v -bench -benchmem ./

get:
	go mod tidy
	go get ./...

gocritic:
	gocritic check ./...

sleep:
	sleep 1

default: get build run sleep test