package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func MakeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ServeOnly("/", ServeOne(INDEX_LOC, "text/html; charset=utf-8")))
	mux.HandleFunc(POST_PREFIX, ServeHTML)
	mux.HandleFunc(CSS_PREFIX, ServeCSS)
	mux.HandleFunc(IMG_PREFIX, ServePNG)
	mux.HandleFunc("/favicon.ico", ServeICO)
	mux.HandleFunc("/img/yd32.ico", ServeICO)
	return mux
}

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	ServeFile(w, r, FindHTML, "text/html; charset=utf-8")
}
func ServePNG(w http.ResponseWriter, r *http.Request) {
	ServeFile(w, r, FindPNG, "image/png")
}
func ServeCSS(w http.ResponseWriter, r *http.Request) {
	ServeFile(w, r, FindCSS, "text/css; charset=utf-8")
}

var ServeICO = ServeOne(ICO_LOC, "image/x-icon")

// ServeFile is used instead of http.ServeContent because I don't need
// mime type sniffing or complicated request type handling
// (my files are not big enough to be worth supporting range requests)
func ServeFile(w http.ResponseWriter, r *http.Request, finder func(r *http.Request) (string, error), cType string) {
	fN, err := finder(r)
	status := http.StatusOK
	if err == NOT_FOUND {
		if fN == "" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		status = http.StatusNotFound
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		LogServerErr("route failure: %s", err)
		return
	}
	f, err := os.Open(fN)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "resource not found", http.StatusNotFound)
			return
		}
		str := fmt.Sprintf("file open failure: %s", err)
		http.Error(w, str, http.StatusInternalServerError)
		LogServerErr(str)
		return
	}
	w.Header().Set("Content-Type", cType)
	w.WriteHeader(status)
	if _, err = io.Copy(w, f); err != nil {
		LogServerErr("io.Copy failure: %s", err)
	}
	f.Close()
}

// ServeOne is for creating handlers that directly use a filepath+mime type
// Basically, if there is no dynamic routing, let's let the
// professional Mux do the routing for me
func ServeOne(name, cType string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		f, err := os.Open(name)
		if err != nil {
			if err == os.ErrNotExist {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			str := fmt.Sprintf("file open failure: %s", err)
			http.Error(w, str, http.StatusInternalServerError)
			LogServerErr(str)
			return
		}
		w.Header().Set("Content-Type", cType)
		if _, err = io.Copy(w, f); err != nil {
			LogServerErr("io.Copy failure: %s", err)
		}
		f.Close()
	}
}

// ServeOnly wraps a http.HandlerFunc with a redirect to exactly one path
func ServeOnly(route string, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != route {
			http.Redirect(w, r, route, http.StatusFound)
			return
		}
		f(w, r)
	}
}
