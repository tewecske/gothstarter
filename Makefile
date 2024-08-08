run: build
	@./bin/app

build:
	@go build -tags dev -o bin/app cmd/app/main.go

templ:
	templ generate --watch --proxy=http://localhost:3000

css:
	tailwindcss -i web/css/app.css -o web/public/styles.css --watch

clean:
	rm -rf tmp/
	rm web/public/styles.css
	find web/templates/ -name '*templ.go' -type f -print0 | xargs -0 rm

