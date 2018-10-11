package main

import (
	"fmt"
	"net/http"
	"strings"
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
	fmt.Fprintln(w, "Hello, my name is ", name)
}

func goodbuy(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprintln(w, "Goodbuy ", name)
}
