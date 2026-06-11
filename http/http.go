package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
