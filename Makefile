APP_NAME = server

GOOSE_DBSTRING ?= "root1:root1234@tcp(localhost:33306)/shopdevgo" # khi Production => "root1:root1234@tcp(mysql_container:3306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema #thư mục chứ các schema.sql cần migration vào database
GOOSE_DRIVER ?= mysql
# =: This is a simple assignment operator. It assigns the value to the variable unconditionally.
# ?=: This is a conditional assignment operator. It assigns the value to the variable only if the variable is not already defined.

# Command cháy ứng dụng: make + tên command

docker_build: 
	docker-compose up -d --build
	docker-compose ps 
docker_stop:
	docker-compose down
dev:         
	go run ./cmd/$(APP_NAME)
docker_up:          
	docker compose -f environment/docker-compose-dev.yaml up

## create new default schema file for migration => goose —dir sql/schema create pre_go_crm_user_c sql 
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one
upse: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse: 
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset
swag: 
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs
sqlcgen:
	sqlc generate
.PHONY: dev run upse downse resetse docker_build docker_stop docker_up

.PHONY: air