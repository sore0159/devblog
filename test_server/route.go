package main

import (
	"errors"
	"net/http"
	"path"
	"path/filepath"
)

const (
	CONTENT_DIR = "resources"
	CSS_PREFIX  = "/css/"
	IMG_PREFIX  = "/img/"
	POST_PREFIX = "/post/"
)

var (
	PAGES_DIR = filepath.Join(CONTENT_DIR, "pages")
	IMG_DIR   = filepath.Join(CONTENT_DIR, "img")
	CSS_DIR   = filepath.Join(CONTENT_DIR, "css")
	ICO_LOC   = filepath.Join(IMG_DIR, "yd32.ico")
	NOT_FOUND = errors.New("resource not found")
	//
	DEFAULT_CSS_LOC = filepath.Join(CSS_DIR, "default.css")
	INDEX_LOC       = filepath.Join(PAGES_DIR, "index.html")
)

func FindHTML(r *http.Request) (filePath string, err error) {
	return filepath.Join(PAGES_DIR, "test.html"), nil
	return "", NOT_FOUND
}
func FindPNG(r *http.Request) (filePath string, err error) {
	return filepath.Join(IMG_DIR, path.Base(r.URL.Path)), nil
	return "", NOT_FOUND
}
func FindCSS(r *http.Request) (filePath string, err error) {
	if r.URL.Path == "/css/default.css" {
		return DEFAULT_CSS_LOC, nil
	}
	return "", NOT_FOUND
}
