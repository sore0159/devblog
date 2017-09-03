package generate

//import t "github.com/sore0159/devblog/generate/templates"

import (
	"html/template"
	"path/filepath"
)

var (
	TMP_INDEX        = TemplateFromFiles("frame", "index", "titlebar", "link_list")
	TMP_MAIN_ARCHIVE = TemplateFromFiles("frame", "main_archive", "titlebar", "link_list")
	TMP_TAG_ARCHIVE  = TemplateFromFiles("frame", "tag_archive", "titlebar", "link_list")
	TMP_POST         = TemplateFromFiles("frame", "body", "titlebar", "link_list", "nav_bar")
	TMP_TEST_POST    = TemplateFromFiles("test_frame", "body", "titlebar", "link_list", "nav_bar")
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

func TemplateFromStrings(tmpls ...string) *template.Template {
	t := template.New("base")
	for _, str := range tmpls {
		if _, err := t.Parse(str); err != nil {
			panic(err)
		}
	}
	return t
}

/*
var (
	TMP_INDEX        = TemplateFromStrings(t.TMP_FRAME, t.TMP_INDEX, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_MAIN_ARCHIVE = TemplateFromStrings(t.TMP_FRAME, t.TMP_MAIN_ARCHIVE, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_TAG_ARCHIVE  = TemplateFromStrings(t.TMP_FRAME, t.TMP_TAG_ARCHIVE, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_POST         = TemplateFromStrings(t.TMP_FRAME, t.TMP_BODY, t.TMP_TITLEBAR, t.TMP_NAVBAR)
)
*/
