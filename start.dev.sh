rm -rf docs

swag init -g cmd/api/main.go --parseDependency --parseInternal

go run cmd/api/main.go