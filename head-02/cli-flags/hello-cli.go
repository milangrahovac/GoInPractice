package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello-cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to",
		},
		cli.BoolFlag{
			Name:   "russian, r",
			Hidden: false,
			Usage:  "Use russian language.",
		},
		cli.BoolFlag{
			Name:   "spanish, s",
			Hidden: false,
			Usage:  "Use spanish language.",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		russian := c.GlobalBool("russian")
		spanish := c.GlobalBool("spanish")

		if russian {
			fmt.Printf("Привет %s!\n", name)
		} else if spanish {
			fmt.Printf("Hola %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
		return nil
	}
	app.Run(os.Args)
}
