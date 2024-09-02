package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/foo/", fooHandler)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
