#!/bin/sh

export MIGRATION_DIR=./migrations
export DB_PORT="5432"
export DB_HOST="localhost"
export DB_NAME="indexer-api"
export DB_USER="postgres"
export DB_SSL=disable
export DB_PASSWORD="qwerty"
export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

if [ "$1" = "--dryrun" ]; then
  goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" status -v
elif [ "$1" = "--down" ]; then
  goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" down -v
elif [ "$1" = "--reset" ]; then
  goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" reset -v
else
  goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v
fi
