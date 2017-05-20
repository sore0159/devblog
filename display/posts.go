package display

import dv "mule/devblog/data"

import (
	"io"
	"path"
	"text/template"
)

const (
	TMP_DIR = "templates"
	TMP_EXT = ".html"
)

func TemplateFrom(names ...string) *template.Template {
	for i, n := range names {
		names[i] = path.Join(TMP_DIR, n+TMP_EXT)
	}
	return template.Must(template.ParseFiles(names...))
}

func FormatPosts(posts []*dv.Data, w io.Writer) error {
	t := TemplateFrom("testing")
	return t.ExecuteTemplate(w, "main", posts)
}
