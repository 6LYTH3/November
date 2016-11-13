package main

import (
	user "./model/user"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func UserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := user.User{
		Name:   "Bob Smith",
		Gender: "male",
		Age:    25,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func main() {
	log.Printf("Come on baby, I'm listening you on port 8080!")

	r := httprouter.New()
	// A handler
	r.GET("/", handler)
	r.GET("/user/:id", UserHandler)
	http.ListenAndServe(":8080", r)
}
