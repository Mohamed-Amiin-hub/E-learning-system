#!/bin/sh

# Wait for Postgres
/wait-for-it.sh db:5432 --timeout=30 --strict -- echo "Postgres is up"

# Wait for Redis
/wait-for-it.sh redis:6379 --timeout=30 --strict -- echo "Redis is up"

# Start the Go app
/elearning

