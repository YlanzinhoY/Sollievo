.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/sollievo
	GOOS=windows GOARCH=amd64 go build -o bin/windows/sollievo.exe
	GOOS=darwin GOARCH=amd64 go build -o bin/macos/sollievo


.PHONY: zip
zip:
	mkdir -p build
	
	cd bin/linux && zip ../../build/sollievo-linux.zip sollievo
	cd bin/windows && zip ../../build/sollievo-windows.zip sollievo.exe
	cd bin/macos && zip ../../build/sollievo-macos.zip sollievo
	cd bin/linux && zip sollievo-linux.zip sollievo
	cd bin/windows && zip sollievo-windows.zip sollievo.exe
	cd bin/macos && zip sollievo-macos.zip sollievo


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