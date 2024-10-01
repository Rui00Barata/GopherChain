package blockchain

import "github.com/rui00barata/GopherChain/internal/block"

type Blockchain struct {
	blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis Block", []byte{})
}

func InitiateChain() *Blockchain {
	return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}
