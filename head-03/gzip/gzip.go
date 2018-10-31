package main

import (
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		compress(file)
	}
}

func compress(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer in.Close()

	out, err := os.Create(filename + ".qz")
	if err != nil {
		return err
	}
}
