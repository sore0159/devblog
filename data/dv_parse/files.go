package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	dv "mule/devblog/data"
	"os"
	"path"
)

const (
	PARSED_DIR = "output"
)

var MISSING_INDEX = errors.New("missing index")

// This should overwrite any previous files that have changed!
func CreateFiles(index *dv.Index, data []*dv.Data) error {
	create := func(n string) (*os.File, error) {
		return os.Create(path.Join(PARSED_DIR, n))
	}
	f, err := create(dv.INDEX_NAME)
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

func GetExisting() (*dv.Index, []*dv.Data, error) {
	files, err := ioutil.ReadDir(PARSED_DIR)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dir: %v", err)
	}
	data := make([]*dv.Data, 0, len(files))
	var index *dv.Index
	for _, fD := range files {
		fName := fD.Name()
		f, err := os.Open(path.Join(PARSED_DIR, fName))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file %s: %s", fD.Name(), err)
		}
		if fName == dv.INDEX_NAME {
			index, err = dv.IndexFromReader(f)
		} else {
			var d *dv.Data
			d, err = dv.DataFromReader(f)
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
