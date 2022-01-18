package main

import (
	"go-web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)
	mux.HandleFunc("/about", handler.AboutHandler)
	mux.HandleFunc("/testMethod", handler.MethodHandler)
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/process", handler.ProcessHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))
	port := ":8080"
	log.Printf("Starting on port %s...", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

