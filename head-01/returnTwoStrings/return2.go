package main

import (
	"fmt"
)

func Names() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {

	n1, n2 := Names()
	fmt.Printf("n1=%s, n2=%s\n", n1, n2)

	n3, _ := Names()
	fmt.Printf("n3=%s\n", n3)
}
