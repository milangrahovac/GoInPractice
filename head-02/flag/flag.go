package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "A name to say hello to.")
var spanish bool
var russian bool

func init() {
	flag.BoolVar(&russian, "russian", false, "Use russian language.")
	flag.BoolVar(&russian, "r", false, "Use russian language.")
	flag.BoolVar(&spanish, "spanish", false, "Use spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use spanish language.")
}

func main() {
	flag.Parse()
	if russian == true {
		fmt.Printf("Привет %s!\n", *name)
	} else if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
