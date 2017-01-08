package route

import (
	controller "../controller"
	user "../model/user"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
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

func Load() *httprouter.Router {
	uc := controller.NewUserController()

	r := httprouter.New()
	// A handler
	r.GET("/", handler)
	r.GET("/user/:id", UserHandler)
	r.GET("/xxx/:id", uc.GetUser)
	r.GET("/adduser/:id", uc.AddUser)
	r.GET("/getalluser", uc.GetAllUser)

	return r
}
