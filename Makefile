ifneq (,$(wildcard .env))
include .env
.EXPORT_ALL_VARIABLES:
endif

COLOR_RESET = \033[0m
COLOR_INFO = \033[32m
COLOR_COMMENT = \033[33m
COLOR_HELP = \033[1;34m
COLOR_BOLD = \033[1m

CONTAINER_APP_NAME = web
PROJECT_NAME = Go Coin Tracker
PROJECT_DESCRIPTION = Simple tracker for cryptocurrency portfolios built with Go, Docker, and PostgreSQL.

SHELL := /bin/bash
CWD := $(shell cd -P -- '$(shell dirname -- "$0")' && pwd -P)
AWS_PROFILE := default
AWS_REPOSITORY := 946241444896.dkr.ecr.eu-west-1.amazonaws.com
UID := $(shell id -u)
GID := $(shell id -g)

.DEFAULT_GOAL := help

##@ Helpers 🚀

.PHONY: help
help: ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "${COLOR_HELP}${PROJECT_NAME}${COLOR_RESET}\n${PROJECT_DESCRIPTION}\n\nUsage:\n make ${COLOR_HELP}<target>${COLOR_RESET}\n"} /^[a-zA-Z_-]+:.*?##/ { printf " ${COLOR_HELP}%-30s${COLOR_RESET} %s\n", $$1, $$2 } /^##@/ { printf "\n${COLOR_BOLD}%s${COLOR_RESET}\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: ## Initialize this project
	docker compose build --build-arg UID=$(UID) --build-arg GID=$(GID) --no-cache

.PHONY: start
start: ## Start this project
	docker compose up --pull always -d --wait

.PHONY: down
down: ## Stop this project
	docker compose down --remove-orphans

.PHONY: bash
bash: ## Takes you inside the container
	docker compose exec $(CONTAINER_APP_NAME) bash

##@ Database 💾

.PHONY: db-create
db-create: ## Create the database, if not exists
	docker compose exec -e PGPASSWORD=$(POSTGRES_PASSWORD) $(CONTAINER_APP_NAME) psql -U $(POSTGRES_USER) -h $(POSTGRES_HOST) -d postgres -c "CREATE DATABASE $(POSTGRES_DB);"

.PHONY: db-drop
db-drop: ## Drop the database
	docker compose exec -e PGPASSWORD=$(POSTGRES_PASSWORD) $(CONTAINER_APP_NAME) psql -U $(POSTGRES_USER) -h $(POSTGRES_HOST) -d postgres -c "DROP DATABASE IF EXISTS $(POSTGRES_DB);"


.PHONY: db-migrate
db-migrate: ## Run database migrations
	docker compose exec $(CONTAINER_APP_NAME) migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DB)?sslmode=disable -path migrations up

.PHONY: db-fresh
db-fresh: ## Drop the database and create a new one with all migrations
db-fresh: db-drop db-create db-migrate
