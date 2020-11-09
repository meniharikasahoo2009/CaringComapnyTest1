package proxyController

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ProxyHandleRequestAndRedirectUserProfile(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Inside ProxyHandle for Authand users")
	username := r.Header.Get("username")

	authurl := fmt.Sprintf("http://localhost:8030/auth/%s", username)
	resp, err := http.Get(authurl)
	fmt.Println("Call made to Auth and response received")

	if err != nil {
		fmt.Println("error in request to Auth")
		err.Error()
	}

	if resp.StatusCode == 200 {
		fmt.Println("returned from Auth and going to get User details")
		userurl := fmt.Sprintf("http://localhost:8040/user/profile/%s", username)
		respo, err := http.Get(userurl)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		fmt.Println(respo)
		json.NewEncoder(w).Encode(respo)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)

}

func ProxyHandleRequestAndRedirectMicroServiceName(w http.ResponseWriter, r *http.Request) {

	urlUsermicroservice := "http://localhost:8040/microservice/name"
	fmt.Println("call to users api to get User details")
	resp, err := http.Get(urlUsermicroservice)

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println(resp)

	json.NewEncoder(w).Encode(resp)
}
