package generate

import t "github.com/sore0159/devblog/generate/templates"

import (
	"html/template"
)

func TemplateFromNames(tmpls ...string) *template.Template {
	t := template.New("base")
	for _, str := range tmpls {
		if _, err := t.Parse(str); err != nil {
			panic(err)
		}
	}
	return t
}

var (
	TMP_INDEX        = TemplateFromNames(t.TMP_FRAME, t.TMP_INDEX, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_MAIN_ARCHIVE = TemplateFromNames(t.TMP_FRAME, t.TMP_MAIN_ARCHIVE, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_TAG_ARCHIVE  = TemplateFromNames(t.TMP_FRAME, t.TMP_TAG_ARCHIVE, t.TMP_TITLEBAR, t.TMP_LINKLIST)
	TMP_POST         = TemplateFromNames(t.TMP_FRAME, t.TMP_BODY, t.TMP_TITLEBAR, t.TMP_NAVBAR)
)
