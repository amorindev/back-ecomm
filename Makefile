include .env

run:
	@go run main.go

compose-dev:
	@docker-compose -f docker-compose.dev.yml --env-file .env up