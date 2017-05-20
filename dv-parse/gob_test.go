package main

import (
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	ix, data, err := GetExisting()
	if err != nil {
		log.Fatal("GET ERROR: ", err)
	}
	log.Println("INDEX: ", *ix)
	log.Println("DATA: ", data)
}
