package groupie

import (
	"net/http"
)

func StyleHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style/style.css")
}
