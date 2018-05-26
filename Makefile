default: run

run: build
	@cd bin; ./sentinel

build: 
	@go build -o bin/sentinel

