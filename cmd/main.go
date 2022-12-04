package main

import (
	"log"
	"tiny-block-chain/blockchain"
	boltdao "tiny-block-chain/dao/bolt"

	"github.com/urfave/cli"
)

func main() {
	bolt := boltdao.New()
	defer func() {
		bolt.Close()
	}()
	blockChain := blockchain.New(bolt)
	if err := blockChain.AddBlock([]byte("A Send 10 BTC to B")); err != nil {
		log.Fatal(err)
	}
	blockChain.PrintBlock()

	//flag := &Flag{}
	app := cli.NewApp()
	app.Name = "tiny-block-chain"
	app.Usage = "tiny block chain cmd工具"
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:        "op",
	// 		Usage:       "Operation for tiny block chain",
	// 		Destination: &flag.Op,
	// 	},
	// }
	app.Commands = []cli.Command{
		{
			Name: "new block chain",
			Action: func(ctx *cli.Context) {

			},
		},
		{
			Name: "add block",
			Action: func(ctx *cli.Context) {
				
			},
		},
		{
			Name: "print tiny block chain",
			Action: func(ctx *cli.Context) error {
				return nil
			},
		},
		{
			Name: "reset bolt db",
			Action: func(ctx *cli.Context) {
				bolt.Reset()
			},
		},
	}
}

type TranscationFlag struct {
}
