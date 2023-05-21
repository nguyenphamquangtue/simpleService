#!bin/bash
export ENV=develop
export SERVER=simpleService
export SECRET_KEY=secret
export PORT=:3000
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=elotus
export POSTGRES_PASSWORD=elotus
export POSTGRES_DB=simpleService

run:
	go run cmd/main.go
