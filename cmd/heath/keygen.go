package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"syscall"

	"github.com/eliothedeman/heath/block"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/urfave/cli.v1"
)

var salt = []byte("a;ljaeyg24nlsjmhasdfuvhasoddf;lkjal")

var keygen = cli.Command{
	Name: "keygen",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "o",
			Value: defaultKeyPath,
			Usage: "Output path for key.",
		},
	},
	Action: kg,
}

func kg(c *cli.Context) {
	pass, err := readPassword()
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(path.Dir(c.String("o")), 0777)
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

	cipherKey := pbkdf2.Key(pass, salt, 4096, 32, sha256.New)

	b, err := aes.NewCipher(cipherKey)
	if err != nil {
		log.Fatal(err)
	}
	b.Encrypt(buff, buff)

	err = ioutil.WriteFile(c.String("o"), buff, 0600)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("New key at %s\n", c.String("o"))

}

func readPassword() ([]byte, error) {
	fmt.Println("Please enter password")
	b1, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}
	fmt.Println("Please renter the same password")
	b2, err := terminal.ReadPassword(syscall.Stdin)
	if err != nil {
		return nil, err
	}

	if bytes.Compare(b1, b2) != 0 {
		fmt.Println("Passwords don't match")
		return readPassword()
	}
	return b1, nil
}
