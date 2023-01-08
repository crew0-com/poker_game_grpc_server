postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb -U root --owner=root crew0_poker

dropdb:
	docer exec -it postgres12 dropdb -U root crew0_poker

create_migration:
	migrate create -ext sql -dir database/migrations -seq $(name)

migrate_up:
	migrate -path database/migrations -database "postgresql://root:secret@localhost:5432/crew0_poker?sslmode=disable" -verbose up

migrate_down:
	migrate	-path database/migrations -database "postgresql://root:secret@localhost:5432/crew0_poker?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migration_create migrate_up migrate_down