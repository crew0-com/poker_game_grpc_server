postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb -U root --owner=root crew0_poker

dropdb:
	docer exec -it postgres12 dropdb -U root crew0_poker

.PHONY: postgres createdb