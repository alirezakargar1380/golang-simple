package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootEndpoint).Methods("GET")
	http.ListenAndServe(":8080", r)
}
