CONCH_VERSION="1.0.0"
CONCH_BUILD_TIME=`date +%s`
CONCH_GIT_REV=`git describe --always --abbrev --dirty --long`

UNAME_S=$(shell uname -s)

FLAGS_PATH=github.com/joyent/conch-shell/pkg/util
BUILD_ARGS = -ldflags="-X ${FLAGS_PATH}.Version=${CONCH_VERSION} -X ${FLAGS_PATH}.BuildTime=${CONCH_BUILD_TIME} -X ${FLAGS_PATH}.GitRev=${CONCH_GIT_REV}"

BUILD = go build ${BUILD_ARGS} 

default: clean vendor check bin/conch ## By default, run 'clean', 'check', 'bin/conch'

first-run: tools ## Install all the dependencies needed to build and test

.PHONY: bin/conch
bin/conch: pkg/**/*.go cmd/conch/*.go vendor ## Build bin/conch
	@echo "==> building bin/conch"
	${BUILD} -o bin/conch cmd/conch/conch.go


clean: ## Remove build products from bin/
	@echo "==> Removing build products from bin/"
	rm -f bin/conch

vendor: ## Install dependencies
	dep ensure -v -vendor-only

update_deps: ## Update dependencies
	dep ensure -update -v

.PHONY: release
release: vendor check ## Build binaries for all supported platforms
	@echo "==> Building for all the platforms"
	GOOS=darwin GOARCH=amd64 ${BUILD} -o release/conch-darwin-amd64 cmd/conch/conch.go
	GOOS=linux GOARCH=amd64 ${BUILD} -o release/conch-linux-amd64 cmd/conch/conch.go
	GOOS=linux GOARCH=arm ${BUILD} -o release/conch-linux-arm cmd/conch/conch.go
	GOOS=solaris GOARCH=amd64 ${BUILD} -o release/conch-solaris-amd64 cmd/conch/conch.go
	GOOS=freebsd GOARCH=amd64 ${BUILD} -o release/conch-freebsd-amd64 cmd/conch/conch.go
	GOOS=openbsd GOARCH=amd64 ${BUILD} -o release/conch-openbsd-amd64 cmd/conch/conch.go

# gem install github_changelog_generator
changelog:
	github_changelog_generator -u joyent -p conch-shell #--due-tag ${CONCH_VERSION}

.PHONY: check
check: ## Ensure that code matchs best practices
	@echo "==> Checking for best practices"
	gometalinter \
		--deadline 10m \
		--vendor \
		--sort="path" \
		--aggregate \
		--enable-gc \
		--disable-all \
		--enable vet \
		--enable deadcode \
		--enable varcheck \
		--enable ineffassign \
		--enable golint \
		--enable gofmt \
		./...
#		--enable goimports \
#		--enable misspell \


tools: ## Download and install all dev/code tools
	@echo "==> Installing dev tools"
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install


.PHONY: help
help: ## Display this help message
	@echo "GNU make(1) targets:"
	@grep -E '^[a-zA-Z_.-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

