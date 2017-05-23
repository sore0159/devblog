package format

import "mule/devblog/parse"

import (
	"io"
	"path/filepath"
	"text/template"
)

const (
	TMP_DIR = "resources"
	TMP_EXT = ".html"
)

func TemplateFrom(names ...string) *template.Template {
	for i, n := range names {
		names[i] = filepath.Join(TMP_DIR, n+TMP_EXT)
	}
	return template.Must(template.ParseFiles(names...))
}

func FormatPosts(posts []*parse.ParsedFile, w io.Writer) error {
	t := TemplateFrom("testing")
	return t.ExecuteTemplate(w, "main", posts)
}
