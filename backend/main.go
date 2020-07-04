package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"io"
	"io/ioutil"

	"encoding/json"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/upload", UploadHandler).Methods("POST")
	r.HandleFunc("/api/files", ListFiles).Methods("GET")
	r.HandleFunc("/api/file/{file}", func(res http.ResponseWriter, req *http.Request) {
		configDir, _ := os.UserConfigDir()
		file, err := os.Open(path.Join(configDir, "upload-scheduler", "files", mux.Vars(req)["file"]))
		if err != nil {
			res.WriteHeader(404)
			res.Write([]byte("404"))
			return
		}
		defer file.Close()
		res.Header().Add("Content-Type", "image/png")
		io.Copy(res, file)
	}).Methods("GET")

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	log.Fatal(server.ListenAndServe())
	fmt.Println("App started")
}

// UploadHandler handles uploads :)
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Access-Control-Allow-Origin", "*")

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

	_, err = io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}
	res.Write([]byte("OK"))
}

func ListFiles(res http.ResponseWriter, req *http.Request) {
	configDir, _ := os.UserConfigDir()
	files, err := ioutil.ReadDir(path.Join(configDir, "upload-scheduler", "files"))

	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, v := range files {
		fileNames = append(fileNames, v.Name())
	}

	resp, _ := json.Marshal(fileNames)

	res.Header().Add("Content-Type", "application/json")
	res.Write(resp)
}
