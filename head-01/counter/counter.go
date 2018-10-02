package main

import (
	"fmt"
	"time"
)

func count() {
	for x := 0; x < 5; x++ {
		fmt.Printf("x = %v\n", x)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 4)
}
