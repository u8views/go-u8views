POSTGRES_DSN="postgresql://u8user:u8pass@localhost:5432/u8views?sslmode=disable"

include Makefile.ansible

env-up:
	docker-compose -f docker-compose.yml --env-file .env up -d

restart:
	docker restart go_u8views_app

logs:
	docker logs go_u8views_app

pg:
	docker exec -it go_u8views_postgres bash

env-down:
	docker-compose -f docker-compose.yml --env-file .env down

env-down-with-clear:
	docker-compose -f docker-compose.yml --env-file .env down --remove-orphans -v # --rmi=all

# make migrate-pgsql-create NAME=init
migrate-pgsql-create:
	# mkdir -p ./internal/storage/schema
	$(eval NAME ?= todo)
	goose -dir ./internal/storage/schema -table schema_migrations postgres $(POSTGRES_DSN) create $(NAME) sql

migrate-pgsql-goose-install:
	docker exec go_u8views_app go install github.com/pressly/goose/v3/cmd/goose@latest
migrate-pgsql-up: migrate-pgsql-goose-install
	docker exec go_u8views_app goose -dir ./internal/storage/schema -table schema_migrations postgres up
migrate-pgsql-redo:
	docker exec go_u8views_app goose -dir ./internal/storage/schema -table schema_migrations postgres redo
migrate-pgsql-down:
	docker exec go_u8views_app goose -dir ./internal/storage/schema -table schema_migrations postgres down
migrate-pgsql-reset:
	docker exec go_u8views_app goose -dir ./internal/storage/schema -table schema_migrations postgres reset
migrate-pgsql-status:
	docker exec go_u8views_app goose -dir ./internal/storage/schema -table schema_migrations postgres status

migrate-all-reset:
	time make migrate-pgsql-reset migrate-pgsql-up

generate-dbs:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

generate-template:
	# go install github.com/valyala/quicktemplate/qtc
	qtc -dir=./internal/templates/v2 -skipLineComments
	git add .

# BENCHTIME=100x make bench
# BENCHTIME=1000x make bench
# BENCHTIME=10000x make bench
bench:
	$(eval BENCHTIME ?= 100x)
	echo "BENCHTIME=$(BENCHTIME) make bench"
	POSTGRES_DSN=$(POSTGRES_DSN) go test ./internal/tests/... -v -bench=. -benchmem -benchtime=$(BENCHTIME)

postgres-fixtures:
	test -f "./console/postgres-fixtures.sql"
	cat ./console/postgres-fixtures.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

postgres-fixtures-count:
	test -f "./console/postgres-fixtures-count.sql"
	cat ./console/postgres-fixtures-count.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

postgres-fixtures-clear:
	test -f "./console/postgres-fixtures-clear.sql"
	cat ./console/postgres-fixtures-clear.sql | docker exec -i go_u8views_postgres psql -d u8views -U u8user

go-mod-update:
	go mod tidy
	go mod vendor

local-go-app-run:
	POSTGRES_DSN=$(POSTGRES_DSN) PORT=:8080 go run ./cmd/v1/main.go

esbuild-minify:
	MINIFY=true npm run --prefix=client esbuild
	tree -h ./public/assets/js

esbuild:
	MINIFY=false npm run --prefix=client esbuild
	tree -h ./public/assets/js

# BIGINT PRIMARY KEY (time, user_id) 1 MONTH * 10 000 = 1.735GB
# BIGINT PRIMARY KEY (time, user_id) 1 YEAR  * 10 000 = 8.447GB
# BIGINT PRIMARY KEY (user_id, time) 1 MONTH * 10 000 = 1.804GB
#    INT PRIMARY KEY (user_id, time) 1 MONTH * 10 000 = 1.804GB
postgres-volume-size:
	docker system df -v | grep go-u8views_postgres-data
	docker stats --no-stream

ssh:
	# cat ~/.ssh/id_rsa.pub | ssh root@45.77.2.17 "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys"
	ssh -t root@45.77.2.17 "cd /var/go/u8views/; bash --login"

ssh-copy-tls-certificates:
	mkdir -p ./.docker/volumes/go/tls-certificates
	scp -r root@45.77.2.17:/var/go/u8views/docker/volumes/go/tls-certificates ./.docker/volumes/go

# POSTGRES_PASSWORD=$(echo "$RANDOM$RANDOM" | md5sum | head -c 16; echo;) make generate-production-environment-file
generate-production-environment-file:
	touch .production.env

	grep -qF 'PORT=' .production.env || echo 'PORT=:80' >> .production.env

	# Database
	grep -qF 'POSTGRES_USER=' .production.env || echo 'POSTGRES_USER="u8user"' >> .production.env
	grep -qF 'POSTGRES_PASSWORD=' .production.env || echo 'POSTGRES_PASSWORD="$(POSTGRES_PASSWORD)"' >> .production.env
	grep -qF 'POSTGRES_DB=' .production.env || echo 'POSTGRES_DB="u8views"' >> .production.env
	grep -qF 'POSTGRES_DSN=' .production.env || echo 'POSTGRES_DSN=postgresql://u8user:$(POSTGRES_PASSWORD)@postgres:5432/u8views?sslmode=disable' >> .production.env

	# OAuth 2.0
	grep -qF 'GITHUB_CLIENT_ID=' .production.env || echo 'GITHUB_CLIENT_ID=' >> .production.env
	grep -qF 'GITHUB_CLIENT_SECRET=' .production.env || echo 'GITHUB_CLIENT_SECRET=' >> .production.env

	cat .production.env
