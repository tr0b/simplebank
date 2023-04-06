-include .env

createdb:
	docker-compose exec postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker-compose exec postgres dropdb ${POSTGRES_DB}
migrateup:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSL_MODE}" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${SSL_MODE}" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tr0b/simplebank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test server mock
