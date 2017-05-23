package main

import (
	"log"
	"mule/devblog"
	"mule/devblog/parse"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	log.Println("TEST TWO")
	index, data, err := GetExisting()
	if err != nil {
		log.Println("Get existing error: ", err)
		return
	}
	log.Printf("Got %d existing\n", len(data))
	for _, d := range data {
		log.Printf("%+v", d)
	}
	oldIndex, err := devblog.GetIndex(PARSED_DIR)
	if err != nil {
		log.Println("Get index error: ", err)
		return
	}
	if len(*oldIndex) != len(*index) {
		log.Println("Index len mismatch: old %d new %d", len(*oldIndex), len(*index))
	}
	for i, d1 := range *oldIndex {
		d2 := (*index)[i]
		log.Printf("Index compared\nold %+v\nnew %+v\n", d1, d2)
	}
}

func TestThree(t *testing.T) {
	log.Println("TEST THREE")
	fN := "tests/test.md"
	d, err := parse.ParseFromFile(fN)
	if err != nil {
		log.Println("Parse error: ", err)
		return
	}
	log.Printf("Parsed %s, %s\n", d.FileName, d.Submitted)
	d2 := new(parse.ParsedFile)
	err = d2.ParseFileName(d.FileName)
	if err != nil {
		log.Println("Error parsing filename: ", err)
		return
	}
	log.Printf("Name Parsed %s, %s\n", d2.FileName, d2.Submitted)
	index := devblog.NewIndex()
	err = index.AddData(&(d2.IndexData))
	if err != nil {
		log.Println("Error adding to index: ", err)
		return
	}
	log.Printf("Index Data %s, %s\n", (*index)[0].FileName, (*index)[0].Submitted)
}
