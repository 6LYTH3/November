package main

import (
	"./route"
	"log"
	"net/http"
)


func main() {
	log.Printf("Come on baby, I'm listening you on port 8080!")

	http.ListenAndServe(":8080", route.Load())
}
