run: build
	@./bin/app

build:
	@go build -o bin/app cmd/app/main.go

css:
	tailwindcss -i views/css/app.css -o public/styles.css --watch


