build:
	@go build -o bin/ecom cmd/main.go

run: build 
	@./bin/ecom

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir cmd/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrations/main.go up


migrate-down:
	@go run cmd/migrations/main.go down 
