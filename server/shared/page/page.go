package page

import (
	"net/http"

	"github.com/PaulioRandall/go-examples/server/shared/logfmt"
)

type PageRenderer interface {
	Render(model interface{}) ([]byte, error)
	PageName() string
}

func Serve(w http.ResponseWriter, renderer PageRenderer, model interface{}) (noErrors bool) {
	return serve(w, renderer, model)
}

func serve(w http.ResponseWriter, renderer PageRenderer, model interface{}) (noErrors bool) {

	html, e := renderer.Render(model)
	if e != nil {
		onHTMLReadError(w, renderer.PageName(), e)
		return false
	}

	e = writeResponse(w, html)
	if e != nil {
		onResponseError(w, e)
		return false
	}

	return true
}

func onHTMLReadError(w http.ResponseWriter, pageName string, e error) {
	logfmt.LogError("Failed to render page '%s', %+v", pageName, e)
	w.WriteHeader(http.StatusInternalServerError)
}

func onResponseError(w http.ResponseWriter, e error) {
	logfmt.LogError("Error writing response body, %+v\n", e)
}

func writeResponse(w http.ResponseWriter, body []byte) error {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, e := w.Write(body)
	return e
}
