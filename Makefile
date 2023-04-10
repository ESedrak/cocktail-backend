postgres:
	docker run --name postgres15alpine -p 5433:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15alpine createdb --username=user --owner=user cocktail_list

dropdb:
	docker exec -it postgres15alpine dropdb cocktail_list

migrateup:
	migrate -path db/migration -database "postgresql://user:secret@localhost:5433/cocktail_list?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://user:secret@localhost:5433/cocktail_list?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc