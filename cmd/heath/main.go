package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/eliothedeman/heath/block"

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
			Action: func(c *cli.Context) {
				err := os.MkdirAll(path.Dir(c.String("o")), 0777)
				if err != nil {
					log.Fatal(err)
				}

				key, kErr := block.GenerateKey()
				if kErr != nil {
					log.Fatal(err)
				}

				buff, eErr := block.MarshalKey(key)
				if eErr != nil {
					log.Fatal(eErr)
				}

				err = ioutil.WriteFile(c.String("o"), buff, 0600)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("New key at %s\n", c.String("o"))
			},
		}}

	app.Run(os.Args)
}
