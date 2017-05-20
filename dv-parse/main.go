// dv-parse takes 1 argument, a file to parse into a post
//
// If the file begins with "TAGS " then everything after that
// on the first line is split by commas into tags for that post
//
// The post is given a submission time of whenever the parse
// command is run
//
// Content is parsed into html via the markdown spec
//
// A unique UID is assigned #TODO
package main

import (
	"fmt"
	"os"
	"time"
)

type Data struct {
	FileName  string
	UID       uint64
	Submitted time.Time
	Tags      []string
	Content   []byte
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " INPUTFILE")
		return
	}
	data, err := Parse(os.Args[1])
	if err != nil {
		fmt.Println("Parse Error: ", err)
		return
	}
	fmt.Println("Parse Success: ", data)
	fmt.Println("Content\n", string(data.Content))
}
