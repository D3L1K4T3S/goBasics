package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
