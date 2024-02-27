build:
	docker build -t chi-rest-api:latest .
generate-docs:
	swag init -g main.go