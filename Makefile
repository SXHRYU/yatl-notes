docker-build-notes:
	docker compose build notes

docker-notes:
	docker compose up -d notes

docker-test:
	docker compose run --rm test

run:
	go run cmd/server/main.go
