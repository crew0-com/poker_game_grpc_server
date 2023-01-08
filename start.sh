#!/bin/sh
set -e

echo "load configurations"
FILE=app.env
source $FILE

echo "run db migrations"
DB_MIGRATION_PATH=database/migrations
MIGRATE_TOOL_PATH=migrate

$MIGRATE_TOOL_PATH -path $DB_MIGRATION_PATH -database "$DB_SOURCE" --verbose up

echo "starting the app"
exec "$@"


