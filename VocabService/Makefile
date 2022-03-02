.DEFAULT_GOAL := run_dev

.PHONY: build
build: 
	go build -o app/build/main/main app/cmd/main/main.go

.PHONY: run_dev
run_dev:
	go run app/cmd/dev/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
