.PHONY: proto
proto:
	protoc -I . --go_out=paths=source_relative:. --micro_out=. api/services/draws/v0/*

.PHONY: docker-service-draws
docker-service-draws:
	@docker-compose -f ./deployments/draws.yaml up -d