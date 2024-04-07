package blockchain

import (
	"bytes"
	"crypto/md5"
)

type BlockChain struct {
	// [*block1, *block2, *block3,...]
	Blocks []*Block
}

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	//getting prev block
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	// creating miner transaction
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}

	//creating new block with data and prev Hash value
	newBlock := CreateBlock(data, prevBlock.Hash, append([]*Transaction{coinbaseTransaction}, transactions...))
	//appending new block to chain
	chain.Blocks = append(chain.Blocks, newBlock)
}

func (b *Block) ComputeHash() {

	//data + prevHash
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})

	//creating hash using md5 algo
	computedHash := md5.Sum(concatenatedData)

	b.Hash = string(computedHash[:])
}
