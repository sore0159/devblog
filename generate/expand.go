package generate

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
)

func Expand(parsed []*ParsedFile) (data []*GeneratedFile, err error) {
	processed := make([]*ProcessedFile, len(parsed))
	for i, p := range parsed {
		processed[i] = Process(p)
	}

	data = make([]*GeneratedFile, len(parsed))
	t := TemplateFrom("frame", "body", "titlebar")
	tags := make(map[string][]*ProcessedFile)
	for i, pf := range processed {
		g, err := GenFile(pf.FileName, t, pf)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data[i] = g
		if pf.NoDate {
			continue
		}
		for _, tg := range pf.ContentTags {
			tags[tg[0]] = append(tags[tg[0]], pf)
		}
	}
	SortByDate(processed)

	g, err := GenFile("archives.html", TemplateFrom("frame", "main_archive", "titlebar", "link_list"), processed)
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %s", err.Error())
	}
	data = append(data, g)

	g, err = GenFile("index.html", TemplateFrom("frame", "index", "titlebar", "link_list"), processed[0])
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %s", err.Error())
	}
	data = append(data, g)

	tList := TemplateFrom("frame", "tag_archive", "titlebar", "link_list")
	for tg, list := range tags {
		pTL := ProcessTaglist(tg, list)
		g, err := GenFile(pTL.FileName, tList, pTL)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data = append(data, g)
	}
	return data, nil
}

const (
	TMP_DIR = "resources"
	TMP_EXT = ".html"
)

func TagLink(tag string) (link string) {
	return "archives_" + strings.Replace(strings.ToLower(tag), " ", "_", -1) + ".html"
}

func TemplateFrom(names ...string) *template.Template {
	for i, n := range names {
		names[i] = filepath.Join(TMP_DIR, n+TMP_EXT)
	}
	return template.Must(template.ParseFiles(names...))
}
