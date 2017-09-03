package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const HTTP_PORT = ":8000"
const DEFAULT_STATIC_DIR = "sore0159.github.io"

func main() {
	dn := make(chan byte)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting server on port", HTTP_PORT)
	// Creating my own server var to have access to server.Shutdown()
	var static_dir string
	if len(os.Args) > 1 {
		if fI, err := os.Stat(os.Args[1]); err != nil && fI.IsDir() {
			static_dir = os.Args[1]
		} else {
			log.Printf("Could not use static dir %s, defaulting to %s\n", os.Args[1], DEFAULT_STATIC_DIR)
			static_dir = DEFAULT_STATIC_DIR
		}
	} else {
		static_dir = DEFAULT_STATIC_DIR
	}
	server := &http.Server{Addr: HTTP_PORT, Handler: MakeMux(static_dir)}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println("Listen and Serve Error:", err)
			dn <- 0
		}
	}()
	select {
	case <-ch:
		fmt.Println("")
		log.Println("Termination signal recieved, stopping server...")
		ctx := context.TODO()
		err := server.Shutdown(ctx)
		if err != nil {
			LogServerErr("shutdown failure: %s", err)
		}
	case <-dn:
		fmt.Println("")
		log.Println("Exiting program...")
	}
}

// Stdout logging may be replaced with file-logging, so
// just creating a simple wrapper func for now
func LogServerErr(str string, args ...interface{}) {
	log.Println(fmt.Errorf(str, args...))
}

func MakeMux(static_dir string) *http.ServeMux {
	mux := http.NewServeMux()
	// const STATIC_DIR = "static"
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, static_dir+"/img/yd32.ico")
	})
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(static_dir+"/img"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(static_dir+"/css"))))
	mux.Handle("/", http.FileServer(http.Dir("generated")))
	//mux.Handle("/", http.FileServer(http.Dir(STATIC_DIR)))
	return mux
}
