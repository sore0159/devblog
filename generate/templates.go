package generate

//import t "github.com/sore0159/devblog/generate/templates"

import (
	"html/template"
	"io"
	"path/filepath"
	tTemp "text/template"
)

var (
	TMP_INDEX        = TemplateFromFiles("frame", "index", "titlebar", "link_list")
	TMP_MAIN_ARCHIVE = TemplateFromFiles("frame", "main_archive", "titlebar", "link_list")
	TMP_TAG_ARCHIVE  = TemplateFromFiles("frame", "tag_archive", "titlebar", "link_list")
	TMP_POST         = TemplateFromFiles("frame", "body", "titlebar", "link_list", "nav_bar")
	TMP_TEST_POST    = TemplateFromFiles("test_frame", "body", "titlebar", "link_list", "nav_bar")
	TMP_RSS          = tTemp.Must(tTemp.ParseFiles(filepath.Join(TMP_DIR, "feed.rss")))
)

const (
	TMP_DIR = "templates"
	TMP_EXT = ".html"
)

func TemplateFromFiles(tmpls ...string) *template.Template {
	for i, str := range tmpls {
		tmpls[i] = filepath.Join(TMP_DIR, str+TMP_EXT)
	}
	return template.Must(template.ParseFiles(tmpls...))
}

type Templator interface {
	ExecuteTemplate(io.Writer, string, interface{}) error
}
