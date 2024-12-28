.PHONY: install build

gen-db-types:
	npx supabase gen types typescript --local > ../app/app/database/db.ts

install:
	@go mod tidy
	@go mod download

build:
	@go build -o bin/server
