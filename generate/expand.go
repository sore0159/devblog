package generate

import (
	"fmt"
	"sort"
	"strings"
)

func Expand(parsed []*ParsedFile) (data []*GeneratedFile, err error) {
	processed := make([]*ProcessedFile, 0, len(parsed))
	posts := make([]*ProcessedFile, 0, len(parsed))
	tags := make(map[string][]*ProcessedFile)
	var indexP *ProcessedFile
	for _, p := range parsed {
		pf := Process(p)
		if indexP == nil && p.FileName == "index" {
			indexP = pf
			continue
		}
		processed = append(processed, pf)
		if pf.NoPost {
			continue
		}
		posts = append(posts, pf)
		for _, tg := range pf.ContentTags {
			tags[tg[0]] = append(tags[tg[0]], pf)
		}
	}
	SortByDate(processed)
	SortByDate(posts)

	data = make([]*GeneratedFile, 0, len(parsed))
	g, err := GenFile("archives.html", TMP_MAIN_ARCHIVE, processed)
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %s", err.Error())
	}
	data = append(data, g)

	g, err = GenFile("index.html", TMP_INDEX, [2]*ProcessedFile{indexP, processed[0]})
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %s", err.Error())
	}
	data = append(data, g)

	tList := TMP_TAG_ARCHIVE
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

	t := TMP_POST
	for _, pf := range processed {
		g, err := GenFile(pf.FileName, t, pf)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data = append(data, g)
	}
	return data, nil
}

func TagLink(tag string) (link string) {
	return "archives_" + strings.Replace(strings.ToLower(tag), " ", "_", -1) + ".html"
}
