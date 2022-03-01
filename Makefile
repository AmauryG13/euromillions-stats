include ./$(BASE_PATH)/scripts/make/variables.mk

.PHONY: build
build:
	@for service in $(EMS_SERVICES); do \
		sh ./$(BASE_PATH)/scripts/build.sh $$service $(BASE_PATH); \
	done

.PHONY: proto
proto:
	@for service in $(EMS_SERVICES); do \
		protoc -I .. --go_out=paths=source_relative:.. api/$$service/$(API_VERSION)/message.proto; \
		protoc -I .. --go_out=paths=source_relative:.. --micro_out=.. api/$$service/$(API_VERSION)/service.proto; \
	done

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker-build
docker-build:
	@for service in $(EMS_SERVICES); do \
		sh ./$(BASE_PATH)/scripts/docker-build.sh $$service $(BASE_PATH); \
	done

.PHONY: deploy-test-service
deploy-test-service:
	@for service in $(EMS_SERVICES); do \
		sh ./$(BASE_PATH)/scripts/docker-deploy-test.sh $$service $(BASE_PATH); \
	done