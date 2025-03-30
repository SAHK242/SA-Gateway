#!/usr/bin/env bash

echo -e "$(printf "Upgrade dependencies...")"
go get -u ./...
go mod tidy
go mod vendor

echo -e "$(printf "===============================")"

echo -e "$(printf "Generating swagger document...")"

swag fmt
swag init -q

echo -e "$(printf "===============================")"

echo -e "$(printf "Compiling...")"
go install

echo -e "$(printf "===============================")"
