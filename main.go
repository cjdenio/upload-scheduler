package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("404 :("))
	})

	r.PathPrefix("/").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	log.Fatal(server.ListenAndServe())
}
