# init global envs
include .env
export

# init OS envs and params
ifeq ($(OS),Windows_NT)
	SHELL = C:/PROGRA~1/Git/bin/bash.exe
	# Если мы на Windows, глобально запрещаем Git Bash коверкать пути Docker
    export MSYS_NO_PATHCONV=1
else
	SHELL := bash
endif

export PROJECT_ROOT=$(CURDIR)

# actions
env-up:
	@docker compose up -d 	todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Clear all environment volumes??? The database is cleared after this operation! [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata && \
		echo "Env files is cleared."; \
	else \
  		echo "Env clearing cancelled"; \
	fi;

env-port-forward:
	@docker compose up -d todoapp-socat-port-forwarder

env-port-close:
	@docker compose down todoapp-socat-port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "Missing required param 'sec'. Example: make migrate-create sec=\"123test\""; \
  		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
	create \
	-ext sql \
	-dir /migrations \
	-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
  		echo "Missing required param \"action\". Example: make migrate-action action=\"down 1\""; \
  		exit 1; \
  	fi; \
	docker compose run --rm todoapp-postgres-migrate \
    	-path /migrations \
    	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
    	"$(action)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down