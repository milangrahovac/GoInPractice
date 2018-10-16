package main

import (
	"net/http"
)

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*?)", goodbye)
	http.ListenAndServe(":8080", rr)
}
