default: build run

run: build
	@sudo ./bin/sentinel

build: 
	@go build -o bin/sentinel

