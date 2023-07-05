export RELEASE_VERSION	?= $(shell git show -q --format=%h)
export DOCKER_REGISTRY	?= ghcr.io/kzaun/moutfire
export BUILD			?= api

all: build start
db-migrate:
	rel migrate
db-rollback:
	rel rollback
gen:
	go generate ./...
build: gen
	go mod tidy
	go build -o bin/api ./cmd/api
test: gen
	go test -race ./...
start:
	export $$(cat .env | grep -v ^\# | xargs) && ./bin/api
	
#下面是编译开发用的
podman:
	podman build -t $(DOCKER_REGISTRY)/$(BUILD):$(RELEASE_VERSION) -f ./build/$(BUILD)/Dockerfile .
push:
	podman push $(DOCKER_REGISTRY)/$(BUILD):$(RELEASE_VERSION)
