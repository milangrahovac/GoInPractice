package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PORT", "8000")
	http.HandleFunc("/", homePage)

	log.Println("Starting service...")
	log.Printf("Listening on the port %s...", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("%v", err)
	}
	fmt.Printf("Server is running on the port %v", port)
	fmt.Println("Server is running on the port")

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "Home Page!")
	fmt.Println("Request!")

}

// go run env_config.go
// or
// PORT=8585 go run env_config.go
