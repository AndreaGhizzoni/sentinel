default: run

run: build
	@cd bin; SSH_GRAPHIC_INPUT=1 ./sentinel shell

build:
	@go build -o bin/sentinel
	@cp settings.json bin