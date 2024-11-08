.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/sollievo
	GOOS=windows GOARCH=amd64 go build -o bin/windows/sollievo.exe
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/sollievo

.PHONY: build-test
build-test:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/sollievo-test

.PHONY: build-branch
build-branch:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/sollievo-branch

.PHONY: build-beta
build-beta:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/sollievo-beta

run-frameworks:
	go run main.go frameworks

run-tools:
	go run main.go tools

run-tests:
	go run main.go tests

run-sqldriver:
	go run main.go sqlDriver

clear-dep:
	go mod tidy