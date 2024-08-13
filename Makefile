run: build
	@./bin/app

.PHONY: build
build:
	@go build -tags dev -o bin/app cmd/app/main.go

.PHONY: templ
templ:
	templ generate --watch --proxy=http://localhost:3000

.PHONY: css
css:
	tailwindcss -i web/css/app.css -o web/public/styles.css --watch

.PHONY: clean
clean:
	rm -rf tmp/
	rm web/public/styles.css
	find web/templates/ -name '*templ.go' -type f -print0 | xargs -0 rm

.PHONY: test
test:
	  go test -race -v -timeout 30s ./...
