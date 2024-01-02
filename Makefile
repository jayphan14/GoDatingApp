postgres:
	docker run --name postgresdating -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
createdb:
	docker exec -it postgresdating createdb --username=root --owner=root datingdb      
dropdb:
	docker exec -it postgresdating dropdb datingdb
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/datingdb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/datingdb?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
