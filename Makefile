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

.PHONY: createdb dropdb migrateup migratedown sqlc
