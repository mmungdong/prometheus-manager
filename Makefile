GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
APPS ?= $(shell ls apps)
path := $(shell pwd)

is_dev := $(shell if [ -d ".git" ]; then echo 1; else echo 0; fi)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: server
server:
	@echo "Starting development server..."
	# 根据is_dev判断是否是开发环境，如果是开发环境，使用config-dev.yaml配置文件，否则使用config-prod.yaml配置文件
	@if [ $(is_dev) -eq 1 ]; then \
		cd ./cmd/prom_server && go run . -config configs/config-dev.yaml; \
	else \
		cd ./cmd/prom_server && go run . -config configs/config-prod.yaml; \
	fi

.PHONY: agent
agent:
	@echo "Starting development agent..."
	# 根据is_dev判断是否是开发环境，如果是开发环境，使用config-dev.yaml配置文件，否则使用config-prod.yaml配置文件
	@if [ $(is_dev) -eq 1 ]; then \
		cd ./cmd/prom_agent && go run . -config configs/config-dev.yaml; \
	else \
		cd ./cmd/prom_agent && go run . -config configs/config-prod.yaml; \
	fi

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: test
# test
test:
	go test -v ./... -cover

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help



TAG ?= latest
REPO ?= docker.hub # TODO: set your repository address

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: docker-build
docker-build: # test ## Build docker image with the manager.
	docker build -t ${REPO}/prometheus-manager:${TAG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${REPO}/prometheus-manager:${TAG}


SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

KUSTOMIZE = $(shell pwd)/bin/kustomize
.PHONY: kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v4@v4.5.2)

.PHONY: deploy-yaml
# Generate deploy yaml.
deploy-yaml: kustomize ## Generate deploy yaml.
	$(call gen-yamls)


PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

define gen-yamls
{\
set -e ;\
[ -f ${PROJECT_DIR}/_output/yamls/build ] || mkdir -p ${PROJECT_DIR}/_output/yamls/build; \
rm -rf ${PROJECT_DIR}/_output/yamls/build/manager; \
cp -rf ${PROJECT_DIR}/config/* ${PROJECT_DIR}/_output/yamls/build/; \
cd ${PROJECT_DIR}/_output/yamls/build/manager; \
${KUSTOMIZE} edit set image controller=${REPO}/prometheus-manager:${TAG}; \
set +x ;\
echo "==== create prometheus-manager.yaml in ${PROJECT_DIR}/_output/yamls/ ====";\
${KUSTOMIZE} build ${PROJECT_DIR}/_output/yamls/build/default > ${PROJECT_DIR}/_output/yamls/prometheus-manager.yaml;\
}
endef