
# file watch doesn't work on WSL :(
# place files to linux fs
templ generate --watch --proxy=http://localhost:3000

# air also polls because of this WSL issue

# Generate CSP sha256 for files
cat <file> | openssl sha256 -binary | openssl base64

