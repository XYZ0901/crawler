package main

import (
	"crawler/frontend/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fsh := http.FileServer(http.Dir("frontend/view"))
	mux.Handle("/static/", http.StripPrefix("/static/", fsh))
	mux.Handle("/search",
		controller.CreateSearchResultHandler("frontend/view/template.html"))
	server := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
