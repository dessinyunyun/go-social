include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down 1

.PHONY: seed
seed:
	@go run cmd/migrate/seed/main.go

.PHONY: dirty-database
dirty-db: 
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) force 7

	