include .env

run:
	@go run cmd/api/main.go

compose-dev:
	@docker-compose -f docker-compose.dev.yml --env-file .env up