package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"triDBMS/api"
)

func main() {
	var PORT string = ":8443"
	if value, has := os.LookupEnv("triDBMSPort"); has {
		PORT = ":" + value
	}

	docs := http.FileServer(http.Dir("./docs"))
	static := http.FileServer(http.Dir("./datas"))

	docs = http.StripPrefix("/docs/", docs)
	static = http.StripPrefix("/static/", static)

	http.Handle("/docs/", docs)
	http.Handle("/static/", static)

	http.HandleFunc("/", api.Redirect)
	http.HandleFunc("/api/", api.Route)

	fmt.Println("TriDBMS is now on https://localhost" + PORT)

	err := http.ListenAndServe(PORT, nil)
	erring(err)
}

func erring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
