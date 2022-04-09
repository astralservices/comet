package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var dir string
	err := godotenv.Load(".env") // dev purposes

	if err != nil {
		log.Fatalf("Some error occured while loading .env file %s", err)
	}

	flag.StringVar(&dir, "dir", os.Getenv("COMET_ROOT"), "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + os.Getenv("COMET_PORT"),

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started at " + "http://127.0.0.1:" + os.Getenv("COMET_PORT"))

	log.Fatal(srv.ListenAndServe())
}
