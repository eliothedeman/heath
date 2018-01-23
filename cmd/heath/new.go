package main

import (
	"gopkg.in/urfave/cli.v1"
)

var newdb = cli.Command{
	Name: "new",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Value: defaultDbPath,
			Usage: "Output path for a the new database",
		},
		cli.StringFlag{
			Name:  "key",
			Value: defaultKeyPath,
			Usage: "Path to the private key to use with this database",
		},
	},
	Action: newDB,
}

func newDB(c *cli.Context) {

}
