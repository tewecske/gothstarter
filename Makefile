run: build
	@./bin/app

build:
	@go build -tags dev -o bin/app cmd/app/main.go

templ:
	templ generate --watch --proxy=http://localhost:3000

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch


