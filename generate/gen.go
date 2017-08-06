package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const GEN_FOLDER = "generated"

func Gen(dir string, dirInfo []os.FileInfo) error {
	parsed := make([]*ParsedFile, len(dirInfo))
	for i, di := range dirInfo {
		if pf, err := ParseFromFile(dir, di); err != nil {
			return fmt.Errorf("parsing failed: %s", err.Error())
		} else {
			parsed[i] = pf
		}
	}
	data := Expand(parsed)
	return Write(dir, data)
}

func Write(dir string, data []*GeneratedFile) error {
	if err := os.Mkdir(filepath.Join(dir, GEN_FOLDER), 0755); err != nil {
		return fmt.Errorf("folder creation failure: %s", err.Error())
	}
	for _, d := range data {
		if err := ioutil.WriteFile(filepath.Join(dir, GEN_FOLDER, d.FileName), []byte(d.Contents), 0644); err != nil {
			return fmt.Errorf("file write failure: %s", err.Error())
		}
	}
	return nil
}
