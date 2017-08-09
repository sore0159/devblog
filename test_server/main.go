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

func main() {
	dn := make(chan byte)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting server on port", HTTP_PORT)
	// Creating my own server var to have access to server.Shutdown()
	server := &http.Server{Addr: HTTP_PORT, Handler: MakeMux()}
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

func MakeMux() *http.ServeMux {
	mux := http.NewServeMux()
	const STATIC_DIR = "sore0159.github.io"
	// const STATIC_DIR = "static"
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, STATIC_DIR+"/img/yd32.ico")
	})
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir(STATIC_DIR+"/img"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(STATIC_DIR+"/css"))))
	mux.Handle("/", http.FileServer(http.Dir("generated")))
	//mux.Handle("/", http.FileServer(http.Dir(STATIC_DIR)))
	return mux
}
