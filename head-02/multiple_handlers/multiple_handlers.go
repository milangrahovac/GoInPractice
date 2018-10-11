package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbuy/", goodbuy)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}
