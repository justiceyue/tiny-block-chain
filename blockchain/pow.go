package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type POW interface {
	Work() ([]byte, int64)
}

const _defaultDifficulty = 16

type POWImpl struct {
	// 待验证的区块
	Block *Block
	// 难度,生成的hash前N为是0
	Difficulty int64
}

func CreatePOWWorker(block *Block) *POWImpl {
	return &POWImpl{Block: block, Difficulty: _defaultDifficulty}
}

func (p *POWImpl) Work() ([]byte, int64) {
	var (
		nonce int64
		hash  [32]byte
	)
	// for {
	// 	h := sha256.New()
	// 	h.Write(p.prepare(nonce))
	// 	hash := hex.EncodeToString(h.Sum(nil))
	// 	difficulty := strings.Repeat("0", int(p.Difficulty))
	// 	fmt.Println(hash, difficulty)
	// 	if strings.HasPrefix(hash, difficulty) {
	// 		break
	// 	}
	// 	nonce++
	// }
	var (
		hashInt big.Int
		target  = big.NewInt(1)
	)
	// 左移256-difficulty位
	// 假设difficulty为1->前1位为0->算出来的hash<2^255
	target = target.Lsh(target, uint(256-p.Difficulty))
	for {
		hash = sha256.Sum256(p.prepare(nonce))
		hashInt.SetBytes(hash[:])
		fmt.Printf("%x \n", hash)
		if target.Cmp(&hashInt) == 1 {
			break
		}
		nonce++
	}
	return hash[:], nonce
}

func (p *POWImpl) prepare(nonce int64) []byte {
	return bytes.Join([][]byte{
		p.Block.PrevBlockHash,
		p.Block.Transactions,
		Int64ToBytes(p.Block.Height),
		Int64ToBytes(nonce),
		Int64ToBytes(p.Block.Timestamp.Unix()),
	}, []byte{})
}
