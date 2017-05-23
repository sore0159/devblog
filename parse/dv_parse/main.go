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
	"mule/devblog/parse"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "[-f][-k][-n] INPUTFILE [MOREFILES...]\n",
			os.Args[0], " -h will print help")
		return
	}
	fNames := make([]string, 0, len(os.Args)-1)
	var keepFlag, forceFlag bool
	for _, test := range os.Args[1:] {
		if test == "-h" {
			PrintHelp()
			return
		} else if test == "-f" {
			forceFlag = true
		} else if test == "-k" {
			keepFlag = true
		} else if test == "-n" {
			fNames = nil
			break
		} else {
			fNames = append(fNames, test)
		}
	}
	index, data, err := GetExisting()
	if err != nil {
		fmt.Printf("Error loading existing data: %s\n", err)
		return
	}
	newData := make([]*parse.ParsedFile, 0, len(fNames))
	handle := func(fName string) error {
		d, err := parse.ParseFromFile(fName)
		if err != nil {
			return err
		}
		for _, id := range *index {
			if d.FileName == id.FileName {
				if forceFlag {
					fmt.Printf("Duplicate filename %s found, -f flag forceing addition)\n", fName)
					break
				} else {
					fmt.Printf("Duplicate filename %s found, skipping (use -f to force)\n", d.FileName)
					return nil
				}
			}
		}
		index.AddData(&(d.IndexData))
		data = append(data, d)
		newData = append(newData, d)
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
	if len(fNames) == 0 {
		fmt.Printf("Data parsed: updating index\n")
	} else {
		fmt.Printf("Data parsed: adding %d new posts and updating index\n", len(newData))
	}
	if err := CreateFiles(index, newData); err != nil {
		fmt.Println("Error creating files: ", err)
		return
	}
	fmt.Println("File creation complete!")
	if keepFlag {
		return
	}
	for _, fN := range fNames {
		err := os.Rename(fN, fN+".parsed")
		if err != nil {
			fmt.Printf("Error moving parsed files: %s\n", err)
			return
		}
	}
}

func PrintHelp() {
	fmt.Println("Usage: ", os.Args[0], "[-h][-f][-k] INPUTFILE [MOREFILES...]\n",
		os.Args[0]+" is a program to parse new devblog files and add them to an existing collection of blogs.  New, parsed (but not templated) files will be created in the directory '"+PARSED_DIR+"'\n",
		"Printing this help message with the -h flag will cause no actions to be run\n",
		"Any number of file names can be added at once, any directory will havi it's contents added (not recursively)\n",
		"Any file whose name is found in the existing index will be skipped unless the -f flag is present\n",
		"If the -k flag is present, parsed files will remain untouched, otherwise they will be have '.parsed' appended to their name.",
	)
}
