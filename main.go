package main

import (
	"fmt"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	mux := mux.NewRouter()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/greets/{name}", greeting)

	fmt.Println("Starting...")
	http.ListenAndServe(":"+port, mux)
}

func greeting(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	name := vars["name"]

	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message":"hello %s"}`, name)
}

func homepage(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-type", "application/json")
	fmt.Fprintf(w, `{"message":"Welcome to the home page!"}`)
}
