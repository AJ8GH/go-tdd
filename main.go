package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aj8gh/gotdd/countdown"
	"github.com/aj8gh/gotdd/greet"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "world")
}

var sleepTime = 1 * time.Second

var sleeper = &countdown.ConfigurableSleeper{Duration: sleepTime, SleepFunc: time.Sleep}

func main() {
	countdown.Countdown(sleeper, os.Stdout)

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
