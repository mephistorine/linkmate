migrations_up:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/linkmate?sslmode=disable" up
migrations_down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/linkmate?sslmode=disable" down
swag_generate:
	swag init --ot json
swag_fmt:
	swag fmt
dev:
	nodemon --exec go run main.go --signal SIGTERM
