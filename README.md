# Gateway Guide

## Prerequisites

- [x] `golang` (version: 1.19+) [Installation Instruction](https://go.dev/learn/)
- [x] `swag` (version: 1.8.6) _Please install the exact version_
    - Install: `go install github.com/swaggo/swag@v1.8.4`

## Run

1. Clone the `.env.example` into `.env` and modify/update all config as your machine
2. Run the following commands

```shell
go get
go run main.go
```

## Swagger

- [x] Documentation: [Link](https://github.com/swaggo/swag)

After running `swag init` and start the application using `go run main.go`. The Swagger UI is available
at `{host}/swagger/index.html` (e.g., `localhost:3000/swagger/index.html`) 