package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aj8gh/gotdd/countdown"
	"github.com/aj8gh/gotdd/greet"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "world")
}

func main() {
	countdown.Countdown(os.Stdout)

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
