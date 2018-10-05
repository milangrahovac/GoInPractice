package main

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}
	err := gcfg.ReadFileInto(&config, "conf.ini")
	if err != nil {
		fmt.Println("Falied to parse config file: %ws.", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
