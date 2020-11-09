package app

import (
	"fmt"
	"net/http"

	"github.com/meniharikasahoo2009/CaringComapnyTest1/SRC/proxyController"
)

func App() {
	fmt.Println("app Started")

	http.HandleFunc("/user/profile", proxyController.ProxyHandleRequestAndRedirectUserProfile)
	http.HandleFunc("/microservice/name", proxyController.ProxyHandleRequestAndRedirectUserProfile)
	http.ListenAndServe(":8080", nil)
}
