include .env
export

docker-build-notes:
	docker compose build notes

docker-notes:
	docker compose up -d notes

docker-test:
	docker compose run --rm test

docker-migrate-up:
	docker compose run --rm migrate up

docker-migrate-down:
	docker compose run --rm migrate down

run:
	go run cmd/server/main.go

format:
	golangci-lint fmt

lint:
	golangci-lint run

# call like `make create-migrations NAME=<name_thing>`
create-migrations:
	migrate create -dir migrations -ext sql -seq $$NAME

# install golang-migrate beforehand 	
# call like `make migrate ACTION="up/down/force [rev]"`
migrate:
	migrate -source file://migrations -database \
	postgres://${POSTGRES__USER}:${POSTGRES__PASSWORD}\
	@localhost:${POSTGRES__PORT}/${POSTGRES__DB}?sslmode=disable \
	$$ACTION
