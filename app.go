package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello Startup Act College!")
	msg := "Hello startup act college!"
	fmt.Fprintf(w, "<h1>%s</h1>", msg)
}

func main() {
	flag.Parse()
	log.Print("Creometry is just getting started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
