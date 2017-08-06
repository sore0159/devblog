package generate

import (
	"html/template"
	"path/filepath"
)

type GeneratedFile struct {
	FileName string
	Contents string
}

func Expand(parsed []*ParsedFile) (data []*GeneratedFile) {
	data = make([]*GeneratedFile, len(parsed))
	for i, p := range parsed {
		g := &GeneratedFile{
			FileName: p.FileName + FINAL_EXT,
			Contents: "TODO:FORMAT\n" + p.Content,
		}
		data[i] = g
	}
	return data
}

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

/*func FormatPosts(data interface{}, w io.Writer) error {
	t := TemplateFrom("testing")
	return t.ExecuteTemplate(w, "main", posts)
}*/
