package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "-1", "port to listen on")
	flag.Parse()
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
