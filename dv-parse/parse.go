package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

var (
	TAGS_PREFIX = []byte("TAGS ")
	TAGS_END    = []byte("\n")
	TAGS_SPLIT  = ","
)

func Parse(fileName string) (*Data, error) {
	unparsed, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	d := new(Data)
	d.FileName = path.Base(fileName)
	d.Submitted = time.Now()
	if bytes.HasPrefix(unparsed, TAGS_PREFIX) {
		var tagLine []byte
		split := bytes.SplitN(unparsed, TAGS_END, 2)
		if len(split) < 2 {
			return nil, fmt.Errorf("Tags split returned %d subslices, need >1", len(split))
		}
		tagLine, unparsed = split[0], split[1]
		tagLine = bytes.TrimPrefix(tagLine, TAGS_PREFIX)
		d.Tags = strings.Split(string(tagLine), TAGS_SPLIT)[0:]
		for i, str := range d.Tags {
			d.Tags[i] = strings.TrimSpace(str)
		}
	}
	d.Content = ParseMarkdown(unparsed)
	// UID assigned when added to Index
	return d, nil
}
