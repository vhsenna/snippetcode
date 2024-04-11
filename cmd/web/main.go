package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "8080", "HTTP network address")
	flag.Parse()

	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))

	r.HandleFunc("/", home)
	r.HandleFunc("/snippet/view", snippetView)
	r.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s\n", *addr)
	if err := http.ListenAndServe(*addr, r); err != nil {
		log.Fatal(err)
	}
}
