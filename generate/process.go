package generate

import (
	"html/template"
	"strings"
	"time"
)

type ProcessedFile struct {
	FileName    string
	Title       string
	ContentTags [][2]string
	Content     template.HTML
	NoDate      bool
	NoTitle     bool
	Published   string
	PubTime     time.Time
	TagNavs     []TagNav
}

func Process(p *ParsedFile) *ProcessedFile {
	pf := &ProcessedFile{
		FileName:    p.FileName + FINAL_EXT,
		Title:       strings.ToUpper(p.Title),
		Content:     p.Content,
		ContentTags: make([][2]string, 0, len(p.Tags)),
		PubTime:     p.Published,
	}
	for _, t := range p.Tags {
		switch t {
		case "NODATE":
			pf.NoDate = true
		case "NOTITLE":
			pf.NoTitle = true
		default:
			pf.ContentTags = append(pf.ContentTags, [2]string{t, TagLink(t)})
		}
	}
	if !pf.NoDate {
		pf.Published = p.Published.Format("Jan 2, 2006 (15:04)")
	}
	return pf
}
