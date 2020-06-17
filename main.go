package main

import (
	"log"
	"net/http"

	"html/template"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("404 :("))
	})

	r.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		index, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Fatal(err)
		}
		index.Execute(res, map[string]string{
			"text": "Hello!",
		})
	})
	r.PathPrefix("/").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./index.html")
	})

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	log.Fatal(server.ListenAndServe())
}
