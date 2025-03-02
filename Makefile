run:
	go run cmd/api/main.go

create-migration:
ifndef name
	$(error name is not set. Usage: make create_migration name=<migration_name>)
endif
	migrate create -seq -ext sql -dir migrations $(name)

migrate-up:
	migrate -path migrations -database "sqlite3://checklists.db" up

migrate-down:
	migrate -path migrations -database "sqlite3://checklists.db" down