.PHONY: all
.DEFAULT_GOAL := help

PROJECT_ROOT:=$(shell git rev-parse --show-toplevel)

# Load env properties , db name, port, etc...
# nb: You can change the default config with `make ENV_CONTEXT=".env.uat" `
ENV_CONTEXT ?= .env.dev
LOCAL_ENV_MINE=.env.mine
ENV_CONTEXT_PATH:=$(PROJECT_ROOT)/$(ENV_CONTEXT)

## Override any default values in the parent .env, with your own
-include $(ENV_CONTEXT_PATH) $(LOCAL_ENV_MINE)

COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO := $(shell basename `git rev-parse --show-toplevel`)
DATE := $(shell date +%Y-%m-%d-%H-%M-%S)



test: ## Run unit tests
	go test -short -cover -failfast ./...

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


#####################
#  Private Targets  #
#####################

log: # log env vars
	@echo "\n=====================================\n"
	@echo "AWS_REGION             $(AWS_REGION)"
	@echo "GOOS                   $(GOOS)"
	@echo "GOARCH                 $(GOARCH)"
	@echo "COMMIT                 $(COMMIT)"
	@echo "BRANCH                 $(BRANCH)"
	@echo "BINARY_BASE_PATH       $(BINARY_BASE_PATH)"
	@echo "AWS_ACCESS_KEY_ID      $(AWS_ACCESS_KEY_ID)"
	@echo "AWS_SECRET_ACCESS_KEY  $(AWS_SECRET_ACCESS_KEY)"
	@echo "AWS_REGION             $(AWS_REGION)"
	@echo "\n====================================\n"
