default: build

run: build
	@sudo ./sentinel

build: 
	@go build

