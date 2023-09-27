
include docker.mk

###############################################################################
# Tooling configuration.
#
# All executable names should be defined as variables so that they can be
# overloaded.
###############################################################################

GO            		?= go
SERVER_MAIN_PACKAGE ?= github.com/alloykh/pack-calc/cmd
SERVER_BINARY 		?= pack-calc

lint:
	golangci-lint run

run:
	@$(GO) run ./cmd/

build:
	@$(GO) build -o $(SERVER_BINARY) $(SERVER_MAIN_PACKAGE)

test:
	@$(GO) test -short ./...


