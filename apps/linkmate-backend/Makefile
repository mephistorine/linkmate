migrations_up:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/linkmate?sslmode=disable" up
migrations_down:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/linkmate?sslmode=disable" down
migrations_status:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/linkmate?sslmode=disable" status
swag_generate:
	swag init --parseInternal --parseFuncBody --generatedTime -o open-api --ot json,go
swag_fmt:
	swag fmt
dev:
	nodemon --ignore ./web --exec go run main.go --signal SIGTERM
build:
	go build main.go
