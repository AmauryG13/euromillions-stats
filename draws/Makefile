GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=pkg/proto --go_out=pkg/proto proto/draws.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o draws cmd/draws/main.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -f ../build/images/Dockerfile.draws -t amauryg13/ems-draws:latest ..