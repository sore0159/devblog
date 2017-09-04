package generate

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const GEN_FOLDER_NAME = "dv_generated"

func Gen(w io.Writer, names []string) error {
	parsed := make([]*ParsedFile, 0, len(names))
	for _, name := range names {
		if pf, err := ParseFromFile(name); err != nil {
			if err == ERR_TIMELESS {
				fmt.Fprintf(w, "Ignored %s\n", name)
			} else {
				return fmt.Errorf("parsing failed: %s", err.Error())
			}
		} else if pf != nil {
			fmt.Fprintf(w, "Parsed %s\n", name)
			parsed = append(parsed, pf)
		} else {
			fmt.Fprintf(w, "Ignored %s\n", name)
		}
	}
	if len(parsed) == 0 {
		return fmt.Errorf("no files parsed")
	}
	data, err := Expand(parsed)
	if err != nil {
		return fmt.Errorf("expanding failed: %s", err.Error())
	}
	return Write(w, data)
}

type GeneratedFile struct {
	FileName string
	Contents string
}

func GenFile(fileName string, t Templator, data interface{}) (*GeneratedFile, error) {
	g := new(GeneratedFile)
	g.FileName = fileName
	b := new(bytes.Buffer)
	if err := t.ExecuteTemplate(b, "frame", data); err != nil {
		return nil, err
	}
	g.Contents = string(b.Bytes())
	return g, nil
}

func Write(w io.Writer, data []*GeneratedFile) error {
	destF := GEN_FOLDER_NAME
	if err := os.Mkdir(destF, 0755); err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("folder creation failure: %s", err.Error())
		} else {
			fmt.Fprintf(w, "Clearing existing folder %s\n", destF)
			if err = os.RemoveAll(destF); err != nil {
				return fmt.Errorf("folder clear failure: %v", err)
			}
			if err = os.Mkdir(destF, 0755); err != nil {
				return fmt.Errorf("folder creation failure: %s", err.Error())
			}
		}
	} else {
		fmt.Fprintf(w, "Creating folder %s\n", destF)
	}
	for _, d := range data {
		if d == nil {
			fmt.Fprintf(w, "Error: nil data\n")
			continue
		}
		fmt.Fprintf(w, "Creating file %s\n", filepath.Join(destF, d.FileName))
		if err := ioutil.WriteFile(filepath.Join(destF, d.FileName), []byte(d.Contents), 0644); err != nil {
			return fmt.Errorf("file write failure: %s", err.Error())
		}
	}
	return nil
}

func TestGen(w io.Writer, name string) (string, error) {
	pf, err := ParseFromFile(name)
	pf.Tags = append(pf.Tags, "TESTING")
	pf.FileName = "TESTING_" + pf.FileName
	if err != nil {
		if err == ERR_TIMELESS {
			pf.Published = time.Now()
		} else {
			return "", fmt.Errorf("parsing failed: %s", err.Error())
		}
	} else if pf == nil {
		return "", fmt.Errorf("parsing failed: %s ignored", name)
	}
	fmt.Fprintf(w, "Parsed %s\n", name)
	pr := Process(pf)
	AddTagNavs("", []*ProcessedFile{pr})
	t := TMP_TEST_POST
	g, err := GenFile(pr.FileName, t, pr)
	if err != nil {
		return "", fmt.Errorf("failed to generate page: %s", err.Error())
	}
	return filepath.Join(GEN_FOLDER_NAME, g.FileName), Write(w, []*GeneratedFile{g})
}
