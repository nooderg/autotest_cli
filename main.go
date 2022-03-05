package main

import (
	"autotest-cli/authentication"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"a"},
			Usage:   "Return token",
			Action: func(c *cli.Context) error {
				fmt.Println(authentication.Login(c.Args()[0], c.Args()[1]))

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
