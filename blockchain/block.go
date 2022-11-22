package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"time"
)

var _genesisHash = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type Block struct {
	Height        int64
	Hash          []byte
	PrevBlockHash []byte
	Transactions  []byte
	Timestamp     time.Time
	Nonce         int64
}

func CreateBlock(tx string, prevBlockHash []byte, height int64) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now(),
		Transactions:  []byte(tx),
	}
	pow := CreatePOWWorker(block)
	block.Hash, block.Nonce = pow.Work()
	return block
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// 生成创世区块
func CreateGenesisBlock(tx string) *Block {
	return CreateBlock(tx, _genesisHash, 0)
}

func (b *Block) Serialize() ([]byte, error) {
	blockBytes := &bytes.Buffer{}
	encoder := gob.NewEncoder(blockBytes)
	if err := encoder.Encode(b); err != nil {
		return nil, err
	}
	return blockBytes.Bytes(), nil
}

func Deserialize(in []byte) (*Block, error) {
	blockBytes := bytes.NewReader(in)
	decoder := gob.NewDecoder(blockBytes)
	block := &Block{}
	if err := decoder.Decode(block); err != nil {
		return nil, err
	}
	return block, nil
}
