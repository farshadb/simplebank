postgres:
	docker run --name=bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_HOST_AUTH_METHOD=trust -d postgres

createdb:
	docker exec -it bank createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:@localhost:5432/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:@localhost:5432/simple_bank?sslmode=disable" -verbose down 

sqlc: 
	sqlc generate
	
test: 
	go test -v -cover -race ./...

.PHONY:	postgres createdb dropdb migrateup migratedown sqlc test