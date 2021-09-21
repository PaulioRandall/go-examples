package govuk

import (
	"net/http"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "private, max-age=259200") // 3 Days
	http.ServeFile(w, r, "webserver"+r.URL.Path)
}
