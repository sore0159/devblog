package generate

// First Google result for 'golang markdown parser'
// Selling points:
//  * Simple API
//  * Been around a while, still worked on
//  * Handles more features than _I_ need
//  * Depends only on golang std-lib
//
// func([]byte) []byte
// parsed := blackfriday.MarkdownCommon(unparsed)
import "github.com/russross/blackfriday"

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	//"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const TIME_FORMAT = "0601021504"
const FINAL_EXT = ".html"

var (
	TAGS_PREFIX = []byte("TAGS: ")
	LINE_SPLIT  = []byte("\n")
	TAGS_SPLIT  = ","
	TIME_ZONE   = time.Now().Location()
)

type ParsedFile struct {
	FileName  string
	Title     string
	Published time.Time
	Tags      []string
	Content   template.HTML
}

func ParseFromFile(name string) (*ParsedFile, error) {
	base := filepath.Base(name)
	ext := filepath.Ext(name)
	if strings.ToLower(ext) != ".md" {
		return nil, nil
	}
	fName := strings.TrimSuffix(base, ext)
	nParts := strings.Split(fName, "_")
	var flag bool
	if len(nParts) < 2 {
		flag = true
	}
	pf := &ParsedFile{
		FileName: fName,
	}

	var err error
	if !flag {
		if pf.Published, err = time.ParseInLocation(TIME_FORMAT, nParts[0], TIME_ZONE); err != nil {
			flag = true
		}
	}
	if flag {
		pf.Title = strings.ToLower(strings.Join(nParts, " "))
	} else {
		pf.Title = strings.ToLower(strings.Join(nParts[1:], " "))
	}
	var content []byte
	if content, err = ioutil.ReadFile(name); err != nil {
		return nil, fmt.Errorf("file %s unreadable: %s", name, err.Error())
	}
	split := bytes.SplitN(content, LINE_SPLIT, 2)
	if len(split) > 1 && bytes.HasPrefix(split[0], TAGS_PREFIX) {
		tagLine := bytes.TrimPrefix(split[0], TAGS_PREFIX)
		pf.Tags = strings.Split(string(tagLine), TAGS_SPLIT)
		for i, str := range pf.Tags {
			pf.Tags[i] = strings.TrimSpace(str)
			if pf.Tags[i] == "NODATE" {
				flag = false
			}
		}
		sort.Strings(pf.Tags)
		content = split[1]
	}
	if flag {
		return nil, nil
	}

	pf.Content = template.HTML(blackfriday.MarkdownCommon(content))
	return pf, nil
}

func (pf ParsedFile) String() string {
	return fmt.Sprintf(`Title: %s

Pub: %s
Tags: %s
Content: %s
`,
		pf.Title,
		pf.Published,
		strings.Join(pf.Tags, ", "),
		pf.Content,
	)
}
