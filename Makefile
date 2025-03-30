# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code,tidy modfile and compile
.PHONY: build
build:
	go get -u ./...
	go mod tidy
	go mod vendor
	swag fmt
	swag init -q
	go install

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-ST1003,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #s