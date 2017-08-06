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
	"os"
	"path/filepath"
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

func ParseFromFile(dir string, f os.FileInfo) (*ParsedFile, error) {
	ext := filepath.Ext(f.Name())
	if strings.ToLower(ext) != ".md" {
		return nil, nil
	}
	fName := strings.TrimSuffix(f.Name(), ext)
	nParts := strings.Split(fName, "_")
	if len(nParts) < 2 {
		return nil, nil //fmt.Errorf("filename to short: %s", f.Name())
	}
	pf := &ParsedFile{
		FileName: fName,
		Title:    strings.ToLower(strings.Join(nParts[1:], " ")),
	}

	var err error
	if pf.Published, err = time.ParseInLocation(TIME_FORMAT, nParts[0], TIME_ZONE); err != nil {
		return nil, nil //fmt.Errorf("filename timestamp unparsable: %s", f.Name())
	}
	var content []byte
	if content, err = ioutil.ReadFile(filepath.Join(dir, f.Name())); err != nil {
		return nil, fmt.Errorf("file %s unreadable: %s", filepath.Join(dir, f.Name()), err.Error())
	}
	split := bytes.SplitN(content, LINE_SPLIT, 2)
	if len(split) > 1 && bytes.HasPrefix(split[0], TAGS_PREFIX) {
		tagLine := bytes.TrimPrefix(split[0], TAGS_PREFIX)
		pf.Tags = strings.Split(string(tagLine), TAGS_SPLIT)
		for i, str := range pf.Tags {
			pf.Tags[i] = strings.TrimSpace(str)
		}
		content = split[1]
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
