postgres:
	docker run --name=bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

mysql:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_PASSWORD=something -d mysql:8.2

createdb:
	docker exec -it bank createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it bank dropdb simplebank

grant:
	docker exec -it bank psql -U root -c "GRANT ALL PRIVILEGES ON DATABASE simplebank TO root;"

revoke:
	docker exec -it bank psql -U root -c "REVOKE ALL PRIVILEGES ON DATABASE simplebank FROM root;"
	
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1
	
sqlc: 
	sqlc generate
	
test: 
	go test -v -cover -race ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/lordfarshad/simplebank/db/sqlc Store

.PHONY:	postgres localpostgres mysql createdb dropdb grant revoke migrateup migrateup1 migratedown migratedown1 sqlc test server mock