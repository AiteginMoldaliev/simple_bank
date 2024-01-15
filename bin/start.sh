#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up


echo "Database migration completed successfully. Starting the application..."
exec "$@"