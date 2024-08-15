Using:
Golang - https://github.com/golang/go - https://go.dev/
Chi - https://github.com/go-chi/chi
Templ - https://github.com/a-h/templ - https://templ.guide/
HTMX - https://github.com/bigskysoftware/htmx - https://htmx.org/
TailwindCSS - https://github.com/tailwindlabs/tailwindcss - https://tailwindcss.com/
SQLX - https://github.com/jmoiron/sqlx - https://jmoiron.github.io/sqlx/
Secure, for CSP mostly - https://github.com/unrolled/secure
SQLite - https://github.com/sqlite/sqlite - https://sqlite.org/

Resources:
https://github.com/anthdm/gothstarter
https://github.com/TomDoesTech/GOTTH




# file watch doesn't work on WSL :(
# place files to linux fs
templ generate --watch --proxy=http://localhost:3000

# air also polls because of this WSL issue

# Generate CSP sha256 for files
cat <file> | openssl sha256 -binary | openssl base64

