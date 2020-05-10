package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"triDBMS/api"
)

func main() {
	var PORT string = ":8443"
	if value, has := os.LookupEnv("triDBMSPort"); has {
		PORT = ":" + value
	}

	mux := &http.ServeMux{}
	srv := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	docs := http.FileServer(http.Dir("./docs"))
	static := http.FileServer(http.Dir("./datas"))

	docs = http.StripPrefix("/docs/", docs)
	static = http.StripPrefix("/static/", static)

	mux.Handle("/docs/", docs)
	mux.Handle("/static/", static)

	mux.HandleFunc("/", api.Redirect)
	mux.HandleFunc("/api/", api.Route)

	fmt.Println("TriDBMS is now on http://localhost" + PORT)

	err := srv.ListenAndServe()
	erring(err)
}

func erring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
