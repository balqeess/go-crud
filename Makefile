# the make postgres command create a postgres15 container and connects itself to the bank 
postgres:
	docker run --name postgres15 --network simple-user -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:15-alpine
createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres testuser

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/testuser?sslmode=disable" -verbose up


.PHONY: postgres createdb  migrateup 