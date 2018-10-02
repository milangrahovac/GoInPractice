package main

import (
	"fmt"
)

func getName() string {
	return "World!"
}

func main() {
	name := getName()
	fmt.Printf("Hello %s\n", name)
}
