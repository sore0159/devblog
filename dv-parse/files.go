package main

// Okay so what this is doing is making a terrible, terrible
// database.  I realize that.  It _works_ and it's not hard
// to replace with a less terrible (i.e. actual) database later.
//
// Possible problems include:
//  * Each post on a page requires a filesystem
//    call.  5 posts on a page, 5 file reads!
//  * An index that is searched just by iterating through an array
//      this means finding all post with a tag means searching ALL
//      (well, most) TAGS OF ALL POSTS
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	PARSED_DIR = "output"
	INDEX_NAME = "index.gob"
)

var MISSING_INDEX = errors.New("missing index")

func GetIndex() (*Index, error) {
	f, err := os.Open(path.Join(PARSED_DIR, INDEX_NAME))
	if err != nil {
		if os.IsNotExist(err) {
			return NewIndex(), nil
		} else {
			return nil, fmt.Errorf("failed to read index file: %v", err)
		}
	}
	defer f.Close()
	index, err := IndexFromReader(f)
	if err != nil {
		return nil, err
	}
	return index, nil
}

// This should overwrite any previous files that have changed!
func CreateFiles(index *Index, data []*Data) error {
	create := func(n string) (*os.File, error) {
		return os.Create(path.Join(PARSED_DIR, n))
	}
	f, err := create(INDEX_NAME)
	if err != nil {
		return err
	}
	err = index.WriteTo(f)
	f.Close()
	if err != nil {
		return err
	}
	for _, d := range data {
		f, err = create(d.GobFileName())
		if err != nil {
			return err
		}
		err = d.WriteTo(f)
		f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// Don't include directory
func (id IndexData) GobFileName() string {
	return fmt.Sprintf("%d.gob", id.UID)
}

func (id IndexData) GetFromDirectory(dir string) (*Data, error) {
	f, err := os.Open(path.Join(dir, id.GobFileName()))
	if err != nil {
		return nil, err
	}
	d, err := DataFromReader(f)
	f.Close()
	return d, err
}

func GetExisting() (*Index, []*Data, error) {
	files, err := ioutil.ReadDir(PARSED_DIR)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dir: %v", err)
	}
	data := make([]*Data, 0, len(files))
	var index *Index
	for _, fD := range files {
		fName := fD.Name()
		f, err := os.Open(path.Join(PARSED_DIR, fName))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file %s: %s", fD.Name(), err)
		}
		if fName == INDEX_NAME {
			index, err = IndexFromReader(f)
		} else {
			var d *Data
			d, err = DataFromReader(f)
			data = append(data, d)
		}
		f.Close()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read data from file %s: %s", fD.Name(), err)
		}
	}
	if len(data) > 0 && index == nil {
		return nil, nil, MISSING_INDEX
	}
	return index, data, nil
}
