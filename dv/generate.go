package main

import (
	"fmt"
	"github.com/sore0159/devblog/generate"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func DvTestGenerate(args []string) error {
	if len(args) != 1 {
		return HELP_ERR
	}
	w := os.Stdout
	// w = ioutil.Discard
	fName, err := generate.TestGen(w, args[0])
	if err != nil {
		return err
	}

	log.Println("Starting browser...")
	cmd := exec.Command("open", fName)
	if err := cmd.Run(); err != nil {
		log.Println("browser run error: ", err)
		fmt.Println("\nPlease start a browser manually and go to " + fName)
	}
	return nil
}
