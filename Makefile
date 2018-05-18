# !make

include $(CURDIR)/.env

TESTPKGS = $(shell go list ./... | grep -v cmd | grep -v test | grep -v vendor | grep -v script | grep -v examples)

REPO := github.com/kelvintaywl/diff
IMAGE_TAG ?= latest


.PHONY: dep
dep:
	dep ensure -v

.PHONY: init
init:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/joho/godotenv/cmd/godotenv
	make dep

.PHONY: test
test:
	godotenv go test -v $(TESTPKGS)

.PHONY: build
build:
	godotenv go build $(REPO)/cmd/diff

.PHONY: docker_build
docker_build:
	godotenv docker build --rm -t kelvintaywl/diff:$(IMAGE_TAG) .

.PHONY: run
run:
	godotenv @echo "TODO: run binary"

.PHONY: docker_run
docker_run:
	godotenv docker run --rm -p 127.0.0.1:$(PORT):9999 -e GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) kelvintaywl/diff:$(IMAGE_TAG)
