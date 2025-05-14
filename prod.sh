#!/bin/bash

echo "Lets get the migrate CLI..."
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

echo "Migrating database..."
migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose force 1 || true
migrate -path ./db/migrations -database "$(DATABASE_URL)" -verbose up

echo "Starting SnapCrumb in Production mode using Render..."
./snapcrumb