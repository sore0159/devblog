package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

type GeneratedFile struct {
	FileName string
	Contents string
}

func Expand(parsed []*ParsedFile) (data []*GeneratedFile, err error) {
	data = make([]*GeneratedFile, len(parsed))
	t := TemplateFrom("frame", "body", "titlebar")
	processed := make([]*ProcessedFile, len(parsed))
	for i, p := range parsed {
		processed[i] = Process(p)
	}

	for i, pf := range processed {
		str, err := pf.CreatePage(t)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		g := &GeneratedFile{
			FileName: pf.FileName,
			Contents: str,
		}
		data[i] = g
	}
	return data, nil
}

type ProcessedFile struct {
	FileName    string
	Title       string
	ContentTags []string
	Content     template.HTML
	NoDate      bool
	NoTitle     bool
	Published   string
}

func Process(p *ParsedFile) *ProcessedFile {
	pf := &ProcessedFile{
		FileName:    p.FileName + FINAL_EXT,
		Title:       strings.ToUpper(p.Title),
		Content:     p.Content,
		ContentTags: make([]string, 0, len(p.Tags)),
	}
	for _, t := range p.Tags {
		switch t {
		case "NODATE":
			pf.NoDate = true
			pf.FileName = strings.Join(strings.Split(pf.FileName, "_")[1:], "_")
		case "NOTITLE":
			pf.NoTitle = true
		default:
			pf.ContentTags = append(pf.ContentTags, t)
		}

	}
	if !pf.NoDate {
		pf.Published = p.Published.Format("Jan 2, 2006 (15:04)")
	}
	return pf
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

func (pf *ProcessedFile) CreatePage(t *template.Template) (string, error) {
	b := new(bytes.Buffer)
	if err := t.ExecuteTemplate(b, "frame", pf); err != nil {
		return "", err
	}
	return string(b.Bytes()), nil
}
