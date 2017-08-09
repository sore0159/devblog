package generate

import (
	"fmt"
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
	g, err := GenFile("archives.html", TMP_MAIN_ARCHIVE, processed)
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %s", err.Error())
	}
	data = append(data, g)

	g, err = GenFile("index.html", TMP_INDEX, processed[0])
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
	for i, pf := range processed {
		g, err := GenFile(pf.FileName, t, pf)
		if err != nil {
			return nil, fmt.Errorf("failed to create page: %s", err.Error())
		}
		data[i] = g
	}
	return data, nil
}

func TagLink(tag string) (link string) {
	return "archives_" + strings.Replace(strings.ToLower(tag), " ", "_", -1) + ".html"
}
