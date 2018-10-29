package main

import (
	"fmt"
)

func main() {
	fmt.Println("Outside goroutine!")
	go func(){
		fmt.Println("Inside goroutine!")
	}

}
