-include .env

createdb:
	docker-compose exec postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker-compose exec postgres dropdb ${POSTGRES_DB}
migrateup1:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_ADDRESS}:${POSTGRES_PORT}/${POSTGRES_DB}" -verbose up 1
migratedown1:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_ADDRESS}:${POSTGRES_PORT}/${POSTGRES_DB}" -verbose down 1
migrateup:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_ADDRESS}:${POSTGRES_PORT}/${POSTGRES_DB}" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DB_SERVER_ADDRESS}:${POSTGRES_PORT}/${POSTGRES_DB}" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/tr0b/simplebank/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
