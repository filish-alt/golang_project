postgres:
docker run --name postgres2 -p 5432:5432 -e POSTGRES_USER=root POSTGRES_PASSWORD=1234 -d postgres:12-alpine

createdb:
docker exec -it postgres2 createdb --username=root --owner=root simpleproject

dropdb:
docker exec -it postgres2 dropdb simpleproject

migrateup:
 migrate -path dbmigrate -database "postgresql://root:1234@192.168.99.100:5432/simpleproject?sslmode=disable" -verbose
 up
migratedown: migrate -path dbmigrate -database "postgresql://root:1234@192.168.99.100:5432/simpleproject?sslmode=disable" -verbose
 down

sqlc:
sqlc generate

-PHONY: cratedb dropdb postgres migrateup migratedown sqlc