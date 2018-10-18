package main

import (
	"net/http"
	"regexp"
)

type regexResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func main() {
	rr := newPathResolver()
	rr.Add("GET /hello", hello)
	rr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*?)", goodbye)
	http.ListenAndServe(":8080", rr)
}

func newPathResolver() *regexResolver {
	return &regexResolver{
		handlers: make(map[string]http.HandleFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

func (r *regexResolver) Add(regex string, handler http.HandlerFunc) {
	r.handlers[regex] = handler
	cache, _ := regexp.Compile[regex]
	r.cache[regex] = cache
}

func (r *regexResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range r.handlers {
		if r.cache[pattern].MatchString(check) == true {
			handlerFunc(res, req)
			return
		}
	}
	http.NotFound(res, req)
}
