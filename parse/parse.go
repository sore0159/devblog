package parse

import (
	"bytes"
	"errors"
	"path"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const TIME_FORMAT = "0601021504"
const FINAL_EXT = ".html"

var (
	TAGS_PREFIX = []byte("TAGS: ")
	LINE_SPLIT  = []byte("\n")
	TAGS_SPLIT  = ","
	TIME_ZONE   = time.Now().Location()
)

func (pf *ParsedFile) ParseFileName(fileName string) error {
	fileName = strings.TrimSuffix(path.Base(fileName), path.Ext(fileName))
	if fileName == "" {
		return errors.New("unparsable file name: blank file name")
	}
	parts := strings.Split(fileName, "_")
	var flag byte
	for _, cr := range parts[0] {
		if unicode.IsNumber(cr) {
			flag = 1
			if tm, err := time.ParseInLocation(TIME_FORMAT, parts[0], TIME_ZONE); err != nil {
				return errors.New("unparsable file name: unparsable time")
			} else {
				pf.Submitted = tm
				flag = 2
			}
		}
		break
	}
	if flag < 2 {
		pf.Submitted = time.Now()
		if flag == 1 {
			parts[0] = pf.Submitted.Format(TIME_FORMAT)
		} else {
			parts = append([]string{pf.Submitted.Format(TIME_FORMAT)}, parts...)
		}
	}
	if len(parts) < 2 {
		return errors.New("unparsable file name: <2 parts")
	}
	title := make([]string, len(parts)-1)
	for i, str := range parts[1:] {
		parts[i+1] = strings.ToLower(str)
		title[i] = strings.ToUpper(str)
	}
	pf.FileName = strings.Join(parts, "_") + FINAL_EXT
	pf.Title = strings.Join(title, " ")
	return nil
}

func (pf *ParsedFile) ParseContent(content []byte) error {
	split := bytes.SplitN(content, LINE_SPLIT, 3)
	var start int
	for i, ln := range split {
		if i == len(split)-1 {
			start = i + 1
			break
		} else if bytes.HasPrefix(ln, TAGS_PREFIX) {
			tagLine := bytes.TrimPrefix(ln, TAGS_PREFIX)
			pf.Tags = strings.Split(string(tagLine), TAGS_SPLIT)
			for i, str := range pf.Tags {
				pf.Tags[i] = strings.TrimSpace(str)
			}
		} else if x, err := strconv.ParseUint(string(ln), 10, 64); err == nil && x != 0 {
			pf.UID = x
		} else {
			start = i + 1
			break
		}
	}
	split = bytes.SplitN(content, LINE_SPLIT, start)
	if len(split) == 0 {
		return errors.New("parsecontent failed on empty content")
	}
	pf.Content = split[len(split)-1]
	return nil
}
