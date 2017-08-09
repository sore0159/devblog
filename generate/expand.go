package generate

import (
	"fmt"
	"html/template"
	"path/filepath"
	"sort"
	"strings"
)

func Expand(parsed []*ParsedFile) (data []*GeneratedFile, err error) {
	processed := make([]*ProcessedFile, len(parsed))
	posts := make([]*ProcessedFile, 0, len(parsed))
	tags := make(map[string][]*ProcessedFile)
	for i, p := range parsed {
		pf := Process(p)
		processed[i] = pf
		if pf.NoDate {
			continue
		}
		posts = append(posts, pf)
		for _, tg := range pf.ContentTags {
			tags[tg[0]] = append(tags[tg[0]], pf)
		}
	}
	SortByDate(processed)
	SortByDate(posts)

	data = make([]*GeneratedFile, len(parsed))
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
	tgList := make([]string, 0, len(tags))
	for tg, list := range tags {
		tgList = append(tgList, tg)
		pTL := ProcessTaglist(tg, list)
		g, err := GenFile(pTL.FileName, tList, pTL)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data = append(data, g)
	}
	sort.Strings(tgList)
	AddTagNavs("", posts)
	for _, tg := range tgList {
		fmt.Println("TAG", tg, " LIST ", tags[tg])
		AddTagNavs(tg, tags[tg])
	}

	t := TemplateFrom("frame", "body", "titlebar", "nav_bar")
	for i, pf := range processed {
		g, err := GenFile(pf.FileName, t, pf)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data[i] = g
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
