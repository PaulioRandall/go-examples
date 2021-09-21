package shared

import (
	"bytes"
	"html/template"
	"net/http"
)

func RenderHTML(w http.ResponseWriter, model interface{}, files ...string) (template.HTML, bool) {

	html, e := renderPage(model, files...)
	if e != nil {
		onHTMLReadError(w, pageName(files), e)
		return "", false
	}

	return template.HTML(html), true
}

func Serve(w http.ResponseWriter, model interface{}, files ...string) (noErrors bool) {

	html, e := renderPage(model, files...)
	if e != nil {
		onHTMLReadError(w, pageName(files), e)
		return false
	}

	e = writeResponse(w, html)
	if e != nil {
		onResponseError(w, e)
		return false
	}

	return true
}

func renderPage(model interface{}, files ...string) ([]byte, error) {
	t, e := template.ParseFiles(files...)
	if e != nil {
		return nil, e
	}
	return executeTemplate(t, model)
}

func executeTemplate(t *template.Template, model interface{}) ([]byte, error) {
	var b bytes.Buffer
	if e := t.Execute(&b, model); e != nil {
		return nil, e
	}
	return b.Bytes(), nil
}

func pageName(files []string) string {
	if len(files) > 0 {
		return files[0]
	}
	return ""
}

func onHTMLReadError(w http.ResponseWriter, pageName string, e error) {
	LogError("Failed to render page '%s', %+v", pageName, e)
	w.WriteHeader(http.StatusInternalServerError)
}

func onResponseError(w http.ResponseWriter, e error) {
	LogError("Error writing response body, %+v\n", e)
}

func writeResponse(w http.ResponseWriter, body []byte) error {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, e := w.Write(body)
	return e
}
