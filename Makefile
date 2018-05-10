COMMIT = $(shell git describe --always)
PLATFORM=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin

default: build

# build generate binary on './bin' directory.
build:
	BUILD_PLATFORMS=$(PLATFORM) BUILD_ARCHS=$(ARCH) ./utils/build-all.sh

lint:
	golint ${GOFILES_NOVENDOR}

vet:
	go vet -v ${GOFILES_NOVENDOR}

test:
	go test -v ${GOFILES_NOVENDOR}

fmt:
	go fmt ./...

release: buildx
	git tag v$(VERSION)
	git push origin v$(VERSION)
	ghr v$(VERSION) releases/v$(VERSION)/

dep:
	dep ensure
	dep status
