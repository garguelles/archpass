.DEFAULT_GOAL := build

RUN_IN_DOCKER = docker compose run api

.PHONY:fmt
fmt:
	$(RUN_IN_DOCKER) go fmt ./...

.PHONY:lint
lint: fmt
	$(RUN_IN_DOCKER) golint ./...

.PHONY:vet
vet: fmt
	$(RUN_IN_DOCKER) go vet ./...

.PHONY:build
build: vet
	$(RUN_IN_DOCKER) go build -o out main.go

.PHONY:tidy
tidy:
	$(RUN_IN_DOCKER) go mod tidy

.PHONY:install
install:
	$(RUN_IN_DOCKER) go get $(package)

.PHONY:vendor
vendor:
	$(RUN_IN_DOCKER) go mod vendor

.PHONY:create_schema
create_schema:
	go run -mod=mod entgo.io/ent/cmd/ent new ${schema}

.PHONY:ent_generate
ent_generate:
	go generate ./ent

.PHONY:apply_migrations
apply_migrations:
	./scripts/apply-migrations.sh

.PHONY:create_migration
create_migration:
	./scripts/create-migration.sh ${name}

.PHONY:create_blank_migrations
create_blank_migrations:
	./scripts/create-blank-migration.sh ${name}

.PHONE:migrate_hash
migrate_hash:
	./scripts/migrate-hash.sh

.PHONY:clean_schema
clean_schema:
	./scripts/clean-schema.sh

.PHONY:inspect_schema
inspect_schema:
	./scripts/inspect-schema.sh
