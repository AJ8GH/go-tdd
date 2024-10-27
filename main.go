package main

import (
	"log"
	"net/http"

	"github.com/aj8gh/gotdd/greet"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
