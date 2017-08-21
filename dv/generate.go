package main

import (
	"fmt"
	"github.com/sore0159/devblog/generate"
	"io/ioutil"
	"os"
	"path/filepath"
)

func DvGenerate(args []string) error {
	var names []string
	if len(args) == 0 {
		args = []string{"."}
	}
	for _, dir := range args {
		dI2, err := ioutil.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("failed to read dir %s:%s", dir, err.Error())
		}
		if len(dI2) == 0 {
			return fmt.Errorf("no files found in %s", dir)
		}
		for _, di := range dI2 {
			names = append(names, filepath.Join(dir, di.Name()))
		}
	}
	w := os.Stdout
	// w = ioutil.Discard
	return generate.Gen(w, names)
}
