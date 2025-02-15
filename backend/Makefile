include .env.local

empty:
	echo "empty"

migrate:
	go run cmd/migration/main.go

seed:
	go run cmd/seed/main.go

rollback:
	go run cmd/rollback/main.go
