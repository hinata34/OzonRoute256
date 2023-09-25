ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=hinata34 password=hinata34 dbname=OZON host=localhost port=5432 sslmode=disable
endif


MIGRATION_FOLDER=$(CURDIR)/migrations
APP_FOLDER=$(CURDIR)/internal/app

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: compose-up
compose-up:
	docker-compose build
	docker-compose up -d postgres

.PHONY: compose-rm
compose-rm:
	docker-compose down

.PHONY: unit-test
unit-tests:
	go test "$(APP_FOLDER)"/*

.PHONY: unit-tests-cover
unit-tests-cover:
	go test -cover "$(APP_FOLDER)"/*

.PHONY: integration-tests
integration-tests: migration-up
	go test -tags=integration $(CURDIR)/tests

.PHONY: integration-tests-down
integration-tests-down: migration-down migration-up
	go test -tags=integration $(CURDIR)/tests