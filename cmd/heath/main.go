package main

import (
	"os"
	"path"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	defaultParentPath = path.Join(os.Getenv("HOME"), ".config/heath")
	defaultKeyPath    = path.Join(defaultParentPath, "private_key")
	defaultDbPath     = path.Join(defaultParentPath, "heath.db")
)

func main() {
	app := cli.NewApp()
	app.Name = "heath"
	app.Author = "Eliot Hedeman"
	app.Usage = "The heath distributed ledger"

	app.Commands = []cli.Command{
		keygen,
		newdb,
	}

	app.Run(os.Args)
}
