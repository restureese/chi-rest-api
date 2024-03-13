build:
	docker build -t chi-rest-api:latest .
generate-docs:
	swag init -g main.go

migrate-db:
	migrate -path migrations -database "postgresql://admin:example@localhost:5432/example_db?sslmode=disable" up