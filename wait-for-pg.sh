#!/bin/bash

host="$1"
shift
cmd="$@"

count=0
max_attempts=30

while ! PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$host" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
  count=$((count + 1))
  if [ $count -ge $max_attempts ]; then
    >&2 echo "Timeout waiting for Postgres"
    exit 1
  fi
done

>&2 echo "Postgres is up - executing command"
echo "Executing command: $cmd"
exec $cmd