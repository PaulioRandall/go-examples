package render

import (
	"bytes"
	"html/template"
)

type pageRenderer struct {
	files []string
}

func NewPageRenderer(files ...string) *pageRenderer {
	return &pageRenderer{
		files: files,
	}
}

func (pr *pageRenderer) Render(model interface{}) ([]byte, error) {
	return Template(model, pr.files...)
}

func (pr *pageRenderer) PageName() string {
	return pr.files[0]
}

func Template(model interface{}, files ...string) ([]byte, error) {
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
