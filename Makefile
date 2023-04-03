include .env

createdb:
	docker-compose exec postgres createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DATABASE}

dropdb:
	docker-compose exec postgres dropdb ${POSTGRES_DATABASE}
migrateup:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${SSL_MODE}" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=${SSL_MODE}" -verbose down

.PHONY: createdb dropdb migrateup migratedown
