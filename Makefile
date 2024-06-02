.PHONY: generate

include .env
export

run-scraper:
	go run cmd/scraper/main.go

migrate-up:
	cd internal/database && migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable -path migrations up

migrate-down:
	cd internal/database && migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable -path migrations down

run-server: 
	go run cmd/api/main.go

generate:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.42
	go get github.com/99designs/gqlgen/internal/imports@v0.17.42
	go get github.com/99designs/gqlgen@v0.17.42
	go run github.com/99designs/gqlgen generate

test:
	go test ./tests

.PHONY: dump-local-db

dump-local-db:
	sudo chown -R $(shell whoami) backups
	docker exec -u $(POSTGRES_USER) $(POSTGRES_CONTAINER) pg_dump -U $(POSTGRES_USER) -d $(POSTGRES_DB) > backups/local_db_backup.sql




