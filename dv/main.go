package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		err = HELP_ERR
	} else if os.Args[1] == "generate" || os.Args[1] == "gen" || os.Args[1] == "g" {
		err = DvGenerate(os.Args[2:])
	} else if os.Args[1] == "test" || os.Args[1] == "t" {
		err = DvTestGenerate(os.Args[2:])
	} else {
		err = DvPublish(os.Stdout, os.Args[1:])
	}

	if err == HELP_ERR {
		PrintHelp(os.Stdout)
	} else if err != nil {
		PrintErr(os.Stderr, err)
	}
}

func PrintHelp(w io.Writer) {
	fmt.Fprint(w,
		`Usage: 
	dv [c] [filenames...]              -- renames[/copies] files with timestamp
	dv t[est] FILENAME                 -- test generation for single file 
	dv g[en[erate]] [directories...]   -- generates static content
 `)
}

func PrintErr(w io.Writer, err error) {
	fmt.Fprintf(w, "There was an error while processing your request:\n%s\n", err.Error())
}

var HELP_ERR = errors.New("Flag error for help message")
