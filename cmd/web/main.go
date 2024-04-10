package main

import (
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", home)
	r.HandleFunc("/snippet/view", snippetView)
	r.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
