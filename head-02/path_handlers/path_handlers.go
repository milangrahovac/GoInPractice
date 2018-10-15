package main

import (
	"fmt"
	"net/http"
	"path"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func main() {
	fmt.Println("hello world")
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbuy*", goodbuy)

}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprintln(res, req)
		}
	}
	http.NotFound(res, req)
}
