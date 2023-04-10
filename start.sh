#!/bin/sh

set -e

echo "run db migrations"
. /app/.env
/app/migrate -path /app/db/migrate -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
