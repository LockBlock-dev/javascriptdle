all: build

build:
	@templ generate view
	@go build -o bin/javascriptdle main.go

templ:
	@templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
