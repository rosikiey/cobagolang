postgresql:
	docker run --name postgresql -e POSTGRES_USER=myusername -e POSTGRES_PASSWORD=mypassword -p 5432:5432  -d postgres 

createddb:
	docker exec -it postgresql createdb --username=myusername --owner=myusername simple_bank

dropdb:
	docker exec -it postgresql dropdb --username=myusername simple_bank 

migrateup:
	migrate -path db/migration -database "postgresql://myusername:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://myusername:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgresql, createddb, dropdb, migrateup, migratedown, sqlc, test