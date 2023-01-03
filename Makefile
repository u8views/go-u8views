POSTGRES_DSN="postgresql://u8user:u8pass@localhost:5432/u8views?sslmode=disable"

up:
	docker-compose up -d

pg:
	docker exec -it go_u8views_postgres bash

down:
	docker-compose down

down-with-clear:
	docker-compose down --remove-orphans -v # --rmi=all

# make migrate-pgsql-create NAME=init
migrate-pgsql-create:
	# mkdir -p ./internal/storage/schema
	$(eval NAME ?= todo)
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) create $(NAME) sql

migrate-pgsql-up:
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) up
migrate-pgsql-redo:
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) redo
migrate-pgsql-down:
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) down
migrate-pgsql-reset:
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) reset
migrate-pgsql-status:
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) status

migrate-all-reset:
	time make migrate-pgsql-reset migrate-pgsql-up

generate-dbs:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

bench:
	go test ./internal/tests/... -v -bench=. -benchmem

postgres-user-fixtures:
	cat ./console/postgres-user-fixtures.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

go-mod-update:
	go get -u
	go mod tidy
	go mod vendor

local-run:
	DSN=$(POSTGRES_DSN) PORT=8080 go run ./cmd/main.go
