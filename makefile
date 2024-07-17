build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/tooling_golang

run-frameworks:
	go run main.go frameworks

run-tools:
	go run main.go tools

run-tests:
	go run main.go tests

run-sqldriver:
	go run main.go sqlDriver