package main

import (
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		compress(file)
	}
}
