package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	user = []string{"Test", "test1"} //can be added more user
)

func main() {
	fmt.Println("main Started for auth")
	routerA := mux.NewRouter()
	routerA.HandleFunc("/auth/{username}", handleRequest)
	http.ListenAndServe(":8030", routerA)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqUser := vars["username"]
	for _, v := range user {
		if v == reqUser {
			fmt.Println("Authenticated in Auth")
			w.Write([]byte{http.StatusOK})
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
}
