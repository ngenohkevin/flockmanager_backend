postgres:
	 docker run --name flocks -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=dylstar20 -d postgres:latest

createdb:
	 docker exec -it flocks createdb --username=root --owner=root flocks_db

dropdb:
	 docker exec -it flocks dropdb flocks_db

migrateup:
	 migrate -path db/migration -database "postgresql://root:dylstar20@localhost:5432/flocks_db?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:dylstar20@localhost:5432/flocks_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test