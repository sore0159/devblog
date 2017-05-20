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
// A unique UID is assigned to each post when it is added to
// a collection of posts, called an Index
package main

import (
	"fmt"
	"io/ioutil"
	dvp "mule/devblog/dv_parse"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "[-f] INPUTFILE [MOREFILES...]\n",
			os.Args[0], " -h will print help")
		return
	}
	fNames := make([]string, 0, len(os.Args)-1)
	var forceFlag bool
	for _, test := range os.Args[1:] {
		if test == "-f" {
			forceFlag = true
		} else if test == "-h" {
			PrintHelp()
			return
		} else {
			fNames = append(fNames, test)
		}
	}
	index, err := dvp.GetIndex(PARSED_DIR)
	if err != nil {
		fmt.Println("Error getting index: ", err)
		return
	}
	if len(*index) == 0 {
		fmt.Println("Empty index fetched: creating new!")
	}
	data := make([]*dvp.Data, 0, len(fNames))
	handle := func(fName string) error {
		d, err := dvp.Parse(fName)
		if err != nil {
			return err
		}
		for _, id := range *index {
			if d.FileName == id.FileName {
				if forceFlag {
					fmt.Printf("Duplicate filename %s found, -f flag forceing addition)\n", fName)
					break
				} else {
					fmt.Printf("Duplicate filename %s found, skipping (use -f to force)\n", fName)
					return nil
				}
			}
		}
		index.AddData(d)
		data = append(data, d)
		return nil
	}
	for _, fName := range fNames {
		fI, err := os.Stat(fName)
		if err != nil {
			fmt.Printf("Error finding file %s: %s\n", fName, err)
			continue
		}
		if fI.IsDir() {
			files, err := ioutil.ReadDir(fName)
			if err != nil {
				fmt.Printf("Error reading dir %s: %s\n", fName, err)
				continue
			}
			for _, fI = range files {
				if err := handle(path.Join(fName, fI.Name())); err != nil {
					fmt.Printf("Error parsing file  %s: %s\n", fI.Name(), err)
				}
			}
		} else {
			if err := handle(fName); err != nil {
				fmt.Printf("Error parsing file  %s: %s\n", fName, err)
			}
		}
	}
	if len(data) == 0 {
		fmt.Println("No data parsed!")
		return
	}
	fmt.Printf("Data parsed: adding %d new posts and updating index\n", len(data))
	if err := CreateFiles(index, data); err != nil {
		fmt.Println("Error creatiing files: ", err)
		return
	}
	fmt.Println("File creation complete!")
}

func PrintHelp() {
	fmt.Println("Usage: ", os.Args[0], "[-h][-f] INPUTFILE [MOREFILES...]\n",
		"Printing this help message with the -h flag will cause no further actions to be run\n",
		"Any number of file names can be added at once, any directory will havi it's contents added (not recursively)\n",
		"Any file whose name is found in the existing index will be skipped unless the -f flag is present")
}
