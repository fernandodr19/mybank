NAME_APP=app
VERSION=dev
OS ?= linux
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_BUILD_TIME=$(shell date '+%Y-%m-%d__%I:%M:%S%p')

.PHONY: test
test:
	@echo "==> Running Tests"
	go test -race -v ./...

.PHONY: compile
compile: clean
	@echo "==> Go Building API"
	@env GOOS=${OS} GOARCH=amd64 go build -v -o  build/${NAME_APP} \
	-ldflags "-X main.BuildGitCommit=$(GIT_COMMIT) -X main.BuildTime=$(GIT_BUILD_TIME)" ./cmd/api/


.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	@GOOS=${OS} go clean -i -x ./...
	@rm -f build/${NAME_APP}
	@rm -f coverage.html
	@rm -f coverage.out

.PHONY: metalint
metalint:

ifeq (, $(shell which $$(go env GOPATH)/bin/golangci-lint))
	@echo "==> installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	go install ./...
	go test -i ./...
endif

	$$(go env GOPATH)/bin/golangci-lint run -c ./.golangci.yml ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests"
	@richgo test -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

