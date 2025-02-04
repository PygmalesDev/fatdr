build:
	@go build -o bin/fatdir main.go

run: build
	@./bin/fatdir