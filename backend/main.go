package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"io"
	"io/ioutil"

	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/cjdenio/upload-scheduler/db"
	"go.mongodb.org/mongo-driver/bson"

	"context"

	"github.com/google/uuid"
)

func main() {
	db.Connect()

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
	r.HandleFunc("/api/file/{file}/delete", func(res http.ResponseWriter, req *http.Request) {
		configDir, _ := os.UserConfigDir()
		if err := os.Remove(path.Join(configDir, "upload-scheduler", "files", mux.Vars(req)["file"])); err != nil {
			res.WriteHeader(500)
			res.Write([]byte(err.Error()))
		} else {
			res.Write([]byte("OK"))
		}
	})

	r.Methods("OPTIONS").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Access-Control-Allow-Origin", "*")
		res.WriteHeader(204)
		res.Write(nil)
	})

	server := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:3000",
	}

	fmt.Println("App started!")
	log.Fatal(server.ListenAndServe())
}

// UploadHandler handles uploads :)
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	fileID := uuid.New().String()

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

	dest, err := os.Create(path.Join(configDir, "upload-scheduler", "files", fileID))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(dest, src)
	if err != nil {
		log.Fatal(err)
	}

	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Write([]byte("OK"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.DB.Database("upload-scheduler").Collection("files").InsertOne(ctx, bson.M{"name": header.Filename, "id": fileID})

	if err != nil {
		log.Fatal(err)
	}
}

// ListFiles is an API method that lists all files
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

	res.Header().Add("Access-Control-Allow-Origin", "*")
	res.Header().Add("Content-Type", "application/json")
	res.Write(resp)
}
