postgres:
	 docker run --name bank2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:14
mysql:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_PASSWORD=postgres -d mysql:8.2

createdb:
	docker exec -it bank2 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it bank2 dropdb simplebank
migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simplebank?sslmode=disable" -verbose down 

sqlc: 
	sqlc generate
	
test: 
	go test -v -cover -race ./...

.PHONY:	postgres mysql createdb dropdb migrateup migratedown sqlc test