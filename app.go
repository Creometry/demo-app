package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello Startup Act College! this is creometry!")
	msg := "Hello startup act college!!! this is creometry!"
	fmt.Fprintf(w, "<h1>%s</h1>", msg)
}

func main() {
	flag.Parse()
	log.Print("Creometry is just getting started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
