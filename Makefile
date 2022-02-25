GO_UNIT_TESTS=./cmd/...
GO=GOPRIVATE=gitlab.com/route GIT_TERMINAL_PROMPT=1 go
GOPATH_BIN=$(subst :,/bin:,$(GOPATH))/bin
export PATH := $(PATH):$(GOPATH_BIN)

GOLINTER_PATH1=$(GOPATH)/bin/golangci-lint
GOLINTER_PATH2=/usr/local/bin/golangci-lint
GOLINTER=$(shell test -x $(GOLINTER_PATH1) && echo $(GOLINTER_PATH1) || (test -x $(GOLINTER_PATH2) && echo $(GOLINTER_PATH2)) || echo "please install GO Linter https://golangci-lint.run/usage/install/")


up:
	docker-compose up --build

down:
	docker-compose down

#test-ci: @ Run tests in gitlab
test-ci: install.gotools
	APP_ENV=ci-test $(MAKE) golinter run-tests

#install.gotools: @ install golang binary necessary
install.gotools:
	@echo ====[Install.GoTools]====================================================================
	$(GO) mod download

#run-tests: @ convenient task to run tests without generate code or start/stop db
run-tests:
	@echo ====[Running Unit Tests]=================================================================
	@mkdir -p coverage
	$(GO) test -v $(GO_UNIT_TESTS) -coverprofile coverage/coverage.fmt fmt
	@$(GO) tool cover -html=coverage/coverage.fmt -o coverage/coverage.html
	@$(GO) tool cover -func coverage/coverage.fmt | grep total:

## Code Quality
#golinter: @ run golinter to check code style
golinter:
	@echo ====[GO LINTER]=============================================================================
	@$(GOLINTER) run

debug:
	air

build:
	go build -o ./tmp/main ./cmd/main.go