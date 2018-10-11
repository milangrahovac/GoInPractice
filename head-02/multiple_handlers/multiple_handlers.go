package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbuy/", goodbuy)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
