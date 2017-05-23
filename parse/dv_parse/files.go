package main

import (
	"fmt"
	"io/ioutil"
	"mule/devblog/parse"
	"mule/devblog/route"
	"path/filepath"
)

const (
	PARSED_DIR = "parsed"
)

// This should overwrite any previous files that have changed!
func CreateFiles(index *route.Index, data []*parse.ParsedFile) error {
	err := route.MakeIndex(PARSED_DIR, index)
	if err != nil {
		return err
	}
	for _, d := range data {
		err := parse.WriteMDToFile(PARSED_DIR, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetExisting() (*route.Index, []*parse.ParsedFile, error) {
	files, err := ioutil.ReadDir(PARSED_DIR)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read dir: %v", err)
	}
	data := make([]*parse.ParsedFile, 0, len(files))
	index := route.NewIndex()
	for _, fD := range files {
		if fD.Name() == route.INDEX_NAME {
			continue
		}
		fName := filepath.Join(PARSED_DIR, fD.Name())
		pf, err := parse.ParseFromFile(fName)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse file %s: %s", fD.Name(), err)
		}
		err = index.AddData(&(pf.IndexData))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to add file %s to index: %s", fD.Name(), err)
		}
		data = append(data, pf)
	}
	return index, data, nil
}
