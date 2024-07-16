package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "-1", "port to listen on")
	flag.Parse()
	tmpl, _ := template.ParseFiles("index.html")
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Text string }{Text: "Testando"})
	})
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
