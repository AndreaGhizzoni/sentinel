default: run

run: build
	@cd bin; ./sentinel shell

build:
	@go build -o bin/sentinel
	@cp settings.json bin