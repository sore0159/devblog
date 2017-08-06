package generate

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const GEN_FOLDER = "generated"

func Gen(w io.Writer, dir string, dirInfo []os.FileInfo) error {
	parsed := make([]*ParsedFile, 0, len(dirInfo))
	for _, di := range dirInfo {
		if pf, err := ParseFromFile(dir, di); err != nil {
			return fmt.Errorf("parsing failed: %s", err.Error())
		} else if pf != nil {
			fmt.Fprintf(w, "Parsed %s\n", di.Name())
			parsed = append(parsed, pf)
		} else {
			fmt.Fprintf(w, "Ignored %s\n", di.Name())
		}
	}
	if len(parsed) == 0 {
		return fmt.Errorf("no files parsed")
	}
	data, err := Expand(parsed)
	if err != nil {
		return fmt.Errorf("expanding failed: %s", err.Error())
	}
	return Write(w, dir, data)
}

type GeneratedFile struct {
	FileName string
	Contents string
}

func GenFile(fileName string, t *template.Template, data interface{}) (*GeneratedFile, error) {
	g := new(GeneratedFile)
	g.FileName = fileName
	b := new(bytes.Buffer)
	if err := t.ExecuteTemplate(b, "frame", data); err != nil {
		return nil, err
	}
	g.Contents = string(b.Bytes())
	return g, nil
}

func Write(w io.Writer, dir string, data []*GeneratedFile) error {
	if err := os.Mkdir(filepath.Join(dir, GEN_FOLDER), 0755); err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("folder creation failure: %s", err.Error())
		} else {
			fmt.Fprintf(w, "Using existing folder %s\n", filepath.Join(dir, GEN_FOLDER))
		}
	} else {
		fmt.Fprintf(w, "Creating folder %s\n", filepath.Join(dir, GEN_FOLDER))
	}
	for _, d := range data {
		if err := ioutil.WriteFile(filepath.Join(dir, GEN_FOLDER, d.FileName), []byte(d.Contents), 0644); err != nil {
			return fmt.Errorf("file write failure: %s", err.Error())
		} else {
			fmt.Fprintf(w, "Creating file %s\n", filepath.Join(dir, GEN_FOLDER, d.FileName))
		}
	}
	return nil
}
