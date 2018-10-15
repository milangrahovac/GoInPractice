package main

import (
	"fmt"
	"net/http"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func main() {
	fmt.Println("hello world")
	pr := newPathResolver()
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}
