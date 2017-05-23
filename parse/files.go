package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Directory and extention are stripped from fileName
// If fileName begins with a number, it will be parsed
// as the submission date/time via const TIME_FORMAT
// If it does not begin with a number, submission time
// is set to time.Now()
//
// Tags are parsed from first line of content if it
// begins with "TAGS: "
//
// UID is parsed from ???
func ParseFromFile(fileName string) (*ParsedFile, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	d := new(ParsedFile)
	err = d.ParseFileName(fileName)
	if err != nil {
		return nil, err
	}
	err = d.ParseContent(content)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// Writes content as a .md file in the specified directory
// with formatted filename, uid(???), and tags
func WriteMDToFile(dir string, data *ParsedFile) error {
	fName := filepath.Join(dir, strings.TrimSuffix(data.FileName, path.Ext(data.FileName))+".md")
	f, err := os.Create(fName)
	if err != nil {
		return err
	}
	defer f.Close()
	if data.UID != 0 {
		if _, err = fmt.Fprintf(f, "%d\n", data.UID); err != nil {
			return err
		}
	}
	if len(data.Tags) > 0 {
		_, err = f.Write(TAGS_PREFIX)
		if err != nil {
			return err
		}
		if _, err = fmt.Fprint(f, strings.Join(data.Tags, " , ")); err != nil {
			return err
		}
		_, err = f.Write(LINE_SPLIT)
		if err != nil {
			return err
		}
	}
	_, err = f.Write(data.Content)
	return err
}
