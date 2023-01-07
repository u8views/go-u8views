POSTGRES_DSN="postgresql://u8user:u8pass@localhost:5432/u8views?sslmode=disable"
TIMESCALEDB_DSN="postgresql://u8user:u8pass@localhost:54321/u8views?sslmode=disable"

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

# make migrate-tssql-create NAME=init
migrate-tssql-create:
	# mkdir -p ./internal/storage-ts/schema
	$(eval NAME ?= todo)
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) create $(NAME) sql

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

migrate-tssql-up:
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) up
migrate-tssql-redo:
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) redo
migrate-tssql-down:
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) down
migrate-tssql-reset:
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) reset
migrate-tssql-status:
	goose -dir ./internal/storage-ts/schema -table schema_migrations postgres $(TIMESCALEDB_DSN) status

migrate-all-reset:
	time make migrate-pgsql-reset migrate-pgsql-up
	time make migrate-tssql-reset migrate-tssql-up

generate-dbs:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

# BENCHTIME=100x make bench
# BENCHTIME=1000x make bench
# BENCHTIME=10000x make bench
bench:
	$(eval BENCHTIME ?= 100x)
	echo "BENCHTIME=$(BENCHTIME) make bench"
	DSN=$(TIMESCALEDB_DSN) go test ./internal/tests/... -v -bench=. -benchmem -benchtime=$(BENCHTIME)

postgres-fixtures:
	test -f "./console/postgres-fixtures.sql"
	cat ./console/postgres-fixtures.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

postgres-fixtures-count:
	test -f "./console/postgres-fixtures-count.sql"
	cat ./console/postgres-fixtures-count.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

postgres-fixtures-clear:
	test -f "./console/postgres-fixtures-clear.sql"
	cat ./console/postgres-fixtures-clear.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

timescaledb-fixtures:
	test -f "./console/postgres-fixtures.sql"
	cat ./console/postgres-fixtures.sql | docker exec -i go_u8views_timescaledb psql -d u8views -U u8user

timescaledb-fixtures-count:
	test -f "./console/postgres-fixtures-count.sql"
	cat ./console/postgres-fixtures-count.sql | docker exec -i go_u8views_timescaledb psql -d u8views -U u8user

timescaledb-fixtures-clear:
	test -f "./console/postgres-fixtures-clear.sql"
	cat ./console/postgres-fixtures-clear.sql | docker exec -i go_u8views_timescaledb psql -d u8views -U u8user

go-mod-update:
	go get -u
	go mod tidy
	go mod vendor

local-run:
	DSN=$(POSTGRES_DSN) PORT=8080 go run ./cmd/main.go

# BIGINT PRIMARY KEY (time, user_id) 1 MONTH * 10 000 = 1.735GB
# BIGINT PRIMARY KEY (time, user_id) 1 YEAR  * 10 000 = 8.447GB
# BIGINT PRIMARY KEY (user_id, time) 1 MONTH * 10 000 = 1.804GB
#    INT PRIMARY KEY (user_id, time) 1 MONTH * 10 000 = 1.804GB
postgres-volume-size:
	docker system df -v | grep go-u8views_postgres-data
	docker stats --no-stream

timescaledb-volume-size:
	docker system df -v | grep go-u8views_timescaledb-data
	docker stats --no-stream
