.PHONY: generate

run-scraper:
	go run cmd/scraper/main.go

migrate-up:
	cd internal/database && migrate -database postgres://postgres:postgres@localhost:5432/movies?sslmode=disable -path migrations up

migrate-down:
	cd internal/database && migrate -database postgres://postgres:postgres@localhost:5432/movies?sslmode=disable -path migrations down

run-server: 
	go run cmd/api/main.go

generate:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.42
	go get github.com/99designs/gqlgen/internal/imports@v0.17.42
	go get github.com/99designs/gqlgen@v0.17.42
	go run github.com/99designs/gqlgen generate

test:
	go test ./tests


