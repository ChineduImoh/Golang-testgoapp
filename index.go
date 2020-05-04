package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello return hello world message
func Hello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "hello, from Azure!\n")
}

func main() {
	router := mux.NewRouter()
	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	router.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
