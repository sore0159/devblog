package main

import dv "mule/devblog/data"
import "mule/devblog/display"

import (
	"log"
	"net/http"
)

const PORT_NUM = ":8080"
const DATA_DIR = "posts"

func main() {
	http.HandleFunc("/", handle)
	log.Println("STARTING SERVER ON PORT ", PORT_NUM)
	err := http.ListenAndServe(PORT_NUM, nil)
	if err != nil {
		log.Println("Server error: ", err)
	}

}

func handle(w http.ResponseWriter, r *http.Request) {
	index, err := dv.GetIndex(DATA_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Index get error: ", err)
		return
	}
	data := make([]*dv.Data, 0, len(*index))
	for _, id := range *index {
		d, err := id.GetFromDirectory(DATA_DIR)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Data get error: ", err)
			return
		}
		data = append(data, d)
	}
	if err := display.FormatPosts(data, w); err != nil {
		log.Println("TEMPLATE ERROR: ", err)
		return
	}
}
