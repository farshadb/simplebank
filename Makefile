postgres:
	 docker run --name bank2 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine

mysql:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_PASSWORD=postgres -d mysql:8.2

createdb:
	docker exec -it bank2 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it bank2 dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down 

sqlc: 
	sqlc generate
	
test: 
	go test -v -cover -race ./...

.PHONY:	postgres mysql createdb dropdb migrateup migratedown sqlc test