include ./env

MIGRATION_PATH = ./db/migration/
DB_URL = postgresql://$(DB_USER):$(DB_PASS)@:%s/%s?sslmode=disable

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq create_$(NAME)_table

migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) up

migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATION_PATH) down