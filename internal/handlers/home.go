package handlers

import (
	"gothstarter/web/templates"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	Render(w, r, templates.Index())
}
