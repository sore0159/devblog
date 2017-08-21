package main

import (
	"errors"
	"fmt"
	"github.com/sore0159/devblog/generate"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func DvPublish(w io.Writer, args []string) error {
	if len(args) == 0 {
		return errors.New("dv publish requires filename commandline arguments")
	}
	var cpy bool
	if args[0] == "c" {
		cpy = true
		args = args[1:]
		if len(args) == 0 {
			return errors.New("dv publish requires filename commandline arguments")
		}
	}

	now := time.Now().Format(generate.TIME_FORMAT)
	errs := []string{}
	for _, fileName := range args {
		if _, err := os.Stat(fileName); err != nil {
			errs = append(errs, err.Error())
			continue
		}

		dir, base := filepath.Split(fileName)
		newName := filepath.Join(dir, now+"_"+base)
		if cpy {
			if err := CopyFile(fileName, newName); err != nil {
				errs = append(errs, fmt.Sprintf("failed to copy file %s: %s", fileName, err.Error()))
			} else if w != nil {
				fmt.Fprintf(w, "copying %s to %s\n", fileName, newName)
			}
		} else {
			if err := os.Rename(fileName, newName); err != nil {
				errs = append(errs, fmt.Sprintf("failed to move file %s: %s", fileName, err.Error()))
			} else if w != nil {
				fmt.Fprintf(w, "renaming %s to %s\n", fileName, newName)
			}
		}

	}
	if len(errs) == len(args) {
		return fmt.Errorf("no files could be read: %s", strings.Join(errs, ", "))
	} else if len(errs) != 0 {
		return fmt.Errorf("%d files failed: %s", len(errs), strings.Join(errs, ", "))
	}
	return nil
}

func CopyFile(source, dest string) error {
	data, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dest, data, 0644)
}
