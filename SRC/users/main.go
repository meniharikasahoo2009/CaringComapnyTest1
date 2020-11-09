package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Usernmae string
	Dob      string
	Age      int
	Email    string
	Phone    string
}

var (
	users = []User{{"Test", "10/10/2010", 10, "testemail@email.com", "+91 000000000"},
		{"test1", "20/20/1995", 15, "testemail2@emailcom", "+0 9381289321"},
	} //can add more user if needed.
)

func main() {
	fmt.Println("main Started for Users")
	routerU := mux.NewRouter()
	routerU.HandleFunc("/microservice/name", microsServiceName)
	routerU.HandleFunc("/user/profile/{username}", userProfile)
	http.ListenAndServe(":8040", routerU)
}
func microsServiceName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Users microservices Microservicename Handler")
	json.NewEncoder(w).Encode("user-microservice")
}
func userProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Users microservices userProfile Handler")
	vars := mux.Vars(r)
	uname := vars["username"]
	for _, v := range users {
		if v.Usernmae == uname {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	json.NewEncoder(w).Encode(http.StatusNotFound)
}
