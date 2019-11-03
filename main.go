package main

import (
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to this site </h1>"))
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":8080", nil)
}
