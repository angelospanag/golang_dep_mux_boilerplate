package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SayHelloWorld says Hello World!
func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", SayHelloWorld)
	http.ListenAndServe(":8080", r)
}
