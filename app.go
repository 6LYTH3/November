package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	log.Printf("Come on baby, I'm listening you on port 8080!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
