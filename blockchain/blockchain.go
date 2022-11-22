package blockchain

import (
	"fmt"
	"sync"

	"tiny-block-chain/storagedriver"

	"github.com/pkg/errors"
)

type BlockChain struct {
	LatestBlockHash []byte
	storage         storagedriver.StorageDriver
}

var (
	defaultBlockChain *BlockChain
	once              = sync.Once{}
)

func New(storage storagedriver.StorageDriver) *BlockChain {
	once.Do(func() {
		genesisBlock := CreateGenesisBlock("genesis block boom~~~")
		defaultBlockChain = &BlockChain{
			LatestBlockHash: genesisBlock.Hash,
			storage:         storage,
		}
		gs, err := genesisBlock.Serialize()
		if err != nil {
			panic(err)
		}
		storage.SetBlock(genesisBlock.Hash, gs)
	})
	return defaultBlockChain
}

func (blc *BlockChain) AddBlock(tx []byte) error {
	if len(blc.LatestBlockHash) == 0 {
		panic(errors.New("nil block chain"))
	}
	value, err := blc.storage.GetBlock(blc.LatestBlockHash)
	if err != nil {
		return err
	}
	preBlock, err := Deserialize(value)
	if err != nil {
		return err
	}
	lastestBlock := CreateBlock(string(tx), preBlock.Hash, preBlock.Height+1)
	ls, err := lastestBlock.Serialize()
	if err != nil {
		return err
	}
	if err := blc.storage.SetBlock(lastestBlock.Hash, ls); err != nil {
		return err
	}
	blc.LatestBlockHash = lastestBlock.Hash
	return nil
}

func (blc *BlockChain) Iter() func() (*Block, bool) {
	currentHash := blc.LatestBlockHash
	return func() (*Block, bool) {
		value, err := blc.storage.GetBlock(currentHash)
		if err != nil {
			return nil, false
		}
		block, err := Deserialize(value)
		currentHash = block.PrevBlockHash
		return block, block.Height > 0
	}
}

func (blc *BlockChain) PrintBlock() {
	iter := blc.Iter()
	for {
		block, hasNext := iter()
		fmt.Print("======================================================\n")
		fmt.Printf("%x \n", block.PrevBlockHash)
		fmt.Printf("%x \n", block.Hash)
		fmt.Printf("%s \n", block.Transactions)
		fmt.Printf("%+v \n", block.Timestamp)
		fmt.Print("======================================================\n")
		if !hasNext {
			break
		}
	}
}
