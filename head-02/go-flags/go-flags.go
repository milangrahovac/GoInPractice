package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Name    string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Russian bool   `short:"r" long:"russian" description:"Use Russian language."`
	Spanish bool   `short:"s" long:"spanish" description:"Use Spanish language."`
}

func main() {
	flags.Parse(&opts)

	if opts.Russian == true {
		fmt.Printf("Привет %s!\n", opts.Name)
	} else if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
