# The binary to build (just the basename).
BIN ?= uploadServer
BASE64_PASSWORD = ZHVtbXlWYWx1ZQ==
IMAGE_NAME = $(DOCKER_REGISTRY)/danifv27/$(BIN)
VCS_USER ?= cmo-sre
VCS_PASSWORD = $(shell echo $(BASE64_PASSWORD) | base64 --decode)
VCS_PROTOCOL ?= https
VCS_URL ?= tools.adidas-group.com/bitbucket/scm/cmodevops/cmo-optimocli.git
ARTIFACTORY_CREDENTIALS ?= $(shell echo cGMtbWF2ZW46RGVjZW1iZXJAMjAxNg==| base64 --decode)

TOP_LEVEL = .

include $(TOP_LEVEL)/common.mk

.PHONY: help
help: ## Show This Help
	@for line in $$(cat ./Makefile $(TOP_LEVEL)/common.mk | grep "##" | grep -v "grep" | sed  "s/:.*##/:/g" | sed "s/\ /!/g"); do verb=$$(echo $$line | cut -d ":" -f 1); desc=$$(echo $$line | cut -d ":" -f 2 | sed "s/!/\ /g"); printf "%-30s--%s\n" "$$verb" "$$desc"; done
