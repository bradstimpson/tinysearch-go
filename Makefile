export GO111MODULE=on

# Check to stop make from complaining on make release v0.0.0
ifeq (release,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
BIN      = $(CURDIR)/bin
M = $(shell printf "\033[34;1m▶\033[0m")
TIMEOUT = 15
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
GIT_STATUS=$(shell git status --porcelain | wc -l | tr -d '[:space:]')
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on go list -f {{.Dir}} ./{cmd,pkg}/...))
TESTPKGS = $(shell env GO111MODULE=on go list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))

## Misc
.PHONY: release
release: ; $(info $(M) releasing…)	@ ## Releasing with version as arg (checks if repo clean)
ifneq ($(GIT_STATUS), 0)
	$(error "There are uncomitted changes - must be clean before release")
endif
	@echo $(RUN_ARGS) > .version
	gorelease

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything in bin/ and tests/
	@rm -rf $(BIN)
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help: ; $(info $(M) This is what you can do with the Makefile…)	@ ## Helping you
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version: ; @ ## Returns the version of the git repository 
	@echo $(VERSION)

.PHONY: lint
lint: ; $(info $(M) running golint…) @ ## Run golint
	golint -set_exit_status $(PKGS)
	golangci-lint run

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	gofmt -l $(PKGS)

# MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
# DATE    ?= $(shell date +%FT%T%z)



## Build All

# .PHONY: all
# all: fmt lint | $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
# 	$Q $(GO) build \
# 		-tags release \
# 		-ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' \
# 		-o $(BIN)/$(basename $(MODULE)) main.go

## Tools

# $(BIN):
# 	@mkdir -p $@
# $(BIN)/%: | $(BIN) ; $(info $(M) building $(PACKAGE)…)
# 	$Q tmp=$$(mktemp -d); \
# 	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
# 		|| ret=$$?; \
# 	   rm -rf $$tmp ; exit $$ret

# GOLINT = $(BIN)/golint
# $(BIN)/golint: PACKAGE=golang.org/x/lint/golint

# GOCOV = $(BIN)/gocov
# $(BIN)/gocov: PACKAGE=github.com/axw/gocov/...

# GOCOVXML = $(BIN)/gocov-xml
# $(BIN)/gocov-xml: PACKAGE=github.com/AlekSi/gocov-xml

# GO2XUNIT = $(BIN)/go2xunit
# $(BIN)/go2xunit: PACKAGE=github.com/tebeka/go2xunit

## Tests
TEST_TARGETS := test-default test-bench test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check tests
test-default: ARGS=-v -cover -tags=integration			## Run verbose with coverage and integration
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. 	## Run benchmarks
test-verbose: ARGS=-v -cover         					## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         						## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): tests
check tests: fmt lint ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	go test ./{cmd,pkg,tests}/... $(ARGS)

.PHONY: test
test: ; $(info $(M) running fast tests…) @ ## Run fast tests no lint fmt integration
	go test ./{cmd,pkg}/... -short

# .PHONY: $(TEST_TARGETS) test-xml check test tests
# test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
# test-short:   ARGS=-short        ## Run only short tests
# test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
# test-race:    ARGS=-race         ## Run tests with race detector
# $(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
# $(TEST_TARGETS): test
# check test tests: fmt lint ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
# 	$Q $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

# test-xml: fmt lint | $(GO2XUNIT) ; $(info $(M) running xUnit tests…) @ ## Run tests with xUnit output
# 	$Q mkdir -p test
# 	$Q 2>&1 $(GO) test -timeout $(TIMEOUT)s -v $(TESTPKGS) | tee test/tests.output
# 	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

# COVERAGE_MODE    = atomic
# COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
# COVERAGE_XML     = $(COVERAGE_DIR)/coverage.xml
# COVERAGE_HTML    = $(COVERAGE_DIR)/index.html
# .PHONY: test-coverage test-coverage-tools
# test-coverage-tools: | $(GOCOV) $(GOCOVXML)
# test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# test-coverage: fmt lint test-coverage-tools ; $(info $(M) running coverage tests…) @ ## Run coverage tests
# 	$Q mkdir -p $(COVERAGE_DIR)
# 	$Q $(GO) test \
# 		-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $(TESTPKGS) | \
# 					grep '^$(MODULE)/' | \
# 					tr '\n' ',' | sed 's/,$$//') \
# 		-covermode=$(COVERAGE_MODE) \
# 		-coverprofile="$(COVERAGE_PROFILE)" $(TESTPKGS)
# 	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
# 	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

