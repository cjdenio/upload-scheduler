package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"io"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/upload", UploadHandler).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	log.Fatal(server.ListenAndServe())
	fmt.Println("App started")
}

// UploadHandler handles uploads :)
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(1000)
	src, header, err := req.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	configDir, err := os.UserConfigDir()

	if _, err := os.Stat(path.Join(configDir, "upload-scheduler", "files")); os.IsNotExist(err) {
		fmt.Println("Creating folder...")
		err := os.MkdirAll(path.Join(configDir, "upload-scheduler", "files"), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	dest, err := os.Create(path.Join(configDir, "upload-scheduler", "files", header.Filename))
	if err != nil {
		log.Fatal(err)
	}

	written, err := io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(written)
	}
	fmt.Println(header.Filename)
	res.Write([]byte("Hello world"))
}
