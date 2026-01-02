package main

import (
	"net/http"
	"log"
	"fmt"
)

func serveHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")	
}

func serveTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Time")
}

func serveHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/hello", serveHello)
	http.HandleFunc("/time", serveTime)
	http.HandleFunc("/health", serveHealth)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
