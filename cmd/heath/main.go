package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	defaultKeyName = os.Getenv("HOME") + "/.config/heath/private_key"
)

func main() {
	app := cli.NewApp()
	app.Name = "heath"
	app.Author = "Eliot Hedeman"
	app.Usage = "The heath distributed ledger"

	app.Commands = []cli.Command{
		{
			Name: "keygen",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "o",
					Value: defaultKeyName,
					Usage: "Output path for key.",
				},
			},
			Action: keygen,
		}}

	app.Run(os.Args)
}
