package main

import (
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"

	"github.com/eliothedeman/cryptio"
	"github.com/eliothedeman/heath/block"
	"github.com/eliothedeman/heath/db"
	"github.com/golang/protobuf/proto"

	"golang.org/x/crypto/pbkdf2"
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
	pass, err := readPassword()
	if err != nil {
		log.Fatalf("Unable to get read password: %s", err.Error())
	}

	aes.NewCipher(pass)
	cipherKey := pbkdf2.Key(pass, salt, 4096, 32, sha256.New)

	b, err := aes.NewCipher(cipherKey)
	if err != nil {
		log.Fatal(err)
	}

	buff, err := ioutil.ReadFile(c.String("key"))
	if err != nil {
		log.Fatal(err)
	}
	b.Decrypt(buff, buff)

	f, err := os.Create(c.String("db"))
	if err != nil {
		log.Fatal(err)
	}

	var key ecdsa.PrivateKey
	var protoKey block.PrivateKey
	err = block.UnmarshalKey(buff, &key)
	if err != nil {
		log.Fatal(err)
	}

	err = proto.Unmarshal(buff, &protoKey)
	if err != nil {
		log.Fatal(err)
	}

	rws := cryptio.ReadWriteSeeker(f, b)
	d := db.NewNonCachedDriver(rws)

	pub, _ := proto.Marshal(protoKey.Public)

	t, err := block.NewTransaction(&key, pub, block.Transaction_PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	blk, err := block.NewBlock(nil, []*block.Transaction{t}, []ecdsa.PublicKey{key.PublicKey})
	if err != nil {
		log.Fatal(err)
	}
	err = d.Write(blk)
	if err != nil {
		log.Fatal(err)
	}
}
