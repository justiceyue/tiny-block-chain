package main

import (
	"log"
	"tiny-block-chain/blockchain"
	boltdao "tiny-block-chain/dao/bolt"
)

func main() {
	bolt := boltdao.New()
	defer func() {
		bolt.Reset()
		bolt.Close()
	}()
	blockChain := blockchain.New(bolt)
	if err := blockChain.AddBlock([]byte("A Send 10 BTC to B")); err != nil {
		log.Fatal(err)
	}
	blockChain.PrintBlock()
}
