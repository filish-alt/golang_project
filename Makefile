postgres:
docker run --name postgres2 -p 5432:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=1234 -d postgres:12-alpine

createdb:
docker exec -it postgres2 createdb --username=root --owner=root simpleproject

dropdb:
docker exec -it postgres2 dropdb simpleproject

migrateup:
 migrate -path db/migration -database "postgresql://root:1234@192.168.99.100:5432/simple_bank?sslmode=disable" -verbose up

migratedown: 
migrate -path db/migration  -database "postgresql://root:1234@192.168.99.100:5432/simple_bank?sslmode=disable" -verbose down

proto:
 protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
dbdocs:
dbdocs build docs/db.dbml
dbschema:
dbml2sql --postgres -o docs/schema.sql docs/db.dbml  
sqlc:
sqlc generate

-PHONY: cratedb dropdb postgres migrateup migratedown sqlc