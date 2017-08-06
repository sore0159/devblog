package main

import (
	"fmt"
	"io/ioutil"
	"mule/devblog/generate"
	"os"
)

func DvGenerate(args []string) error {
	dir := "."
	if len(args) == 1 {
		dir = args[0]
	} else if len(args) != 0 {
		return HELP_ERR
	}
	dirInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read dir %s:%s", dir, err.Error())
	}
	if len(dirInfo) == 0 {
		return fmt.Errorf("no files found in %s", dir)
	}
	w := os.Stdout
	// w = ioutil.Discard
	return generate.Gen(w, dir, dirInfo)
}
