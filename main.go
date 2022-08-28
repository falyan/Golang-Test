package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Service test skill Golang Berhasil")
}
