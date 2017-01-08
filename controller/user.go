package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	user "../model/user"
	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := user.User{
		Name:   "Bob Smith",
		Gender: "male",
		Age:    50,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetAllUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := new(user.User)
	uj, err := json.Marshal(u.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) AddUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Printf("[!] CTR Add User")
	u := user.User{
		Name:   "Blythe",
		Gender: "male",
		Age:    25,
		Id:     p.ByName("id"),
	}

	u.Add()

	uj, err := json.Marshal(u.FindAll())
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
