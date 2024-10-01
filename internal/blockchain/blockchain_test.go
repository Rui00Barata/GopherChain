package blockchain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitiateBlockchain(t *testing.T) {
	assert := assert.New(t)

	bc := InitiateChain()

	assert.NotNil(bc, "Blockchain should not be nil")
	assert.Equal(1, len(bc.blocks), "Blockchain should have one genesis block")
	assert.Equal("Genesis Block", string(bc.blocks[0].Data), "First block should be genesis block")
}

func TestAddBlock(t *testing.T) {
	assert := assert.New(t)

	bc := InitiateChain()
	bc.AddBlock("Test Block")

	assert.Equal(2, len(bc.blocks), "Blockchain should havbe two blocks after adding one")
	assert.Equal("Test Block", string(bc.blocks[1].Data), "Second block should have correct data")
	assert.Equal(bc.blocks[0].Hash, bc.blocks[1].PrevBlockHash, "New block should reference the precious block's hash")
}

func TestBlockchainIntegrity(t *testing.T) {
	assert := assert.New(t)
	bc := InitiateChain()

	blocks := []string{"Block 1", "Block 2", "Block 3"}
	for _, block := range blocks {
		bc.AddBlock(block)
	}

	assert.Equal(len(blocks)+1, len(bc.blocks), "Blockchain should have correct number of blocks")

	for i, block := range bc.blocks[1:] {
		assert.Equal(blocks[i], string(block.Data), "Block data should match input")
		assert.Equal(bc.blocks[i].Hash, block.PrevBlockHash, "Block should reference correct previous hash")
	}
}

func TestBlockchainConsistency(t *testing.T) {
	assert := assert.New(t)
	bc1 := InitiateChain()
	bc2 := InitiateChain()

	blocks := []string{"Block 1", "Block 2", "Block 3"}
	for _, block := range blocks {
		bc1.AddBlock(block)
		bc2.AddBlock(block)
	}

	assert.Equal(len(bc1.blocks), len(bc2.blocks), "Both blockchains should have the same number of blocks")

	for i := range bc1.blocks {
		assert.Equal(bc1.blocks[i].Data, bc2.blocks[i].Data, "Corresponding blocks should have the same data")
	}
}

func TestLargeBlockchain(t *testing.T) {
	assert := assert.New(t)
	bc := InitiateChain()

	for i := 0; i < 1000; i++ {
		bc.AddBlock(fmt.Sprintf("Block %d", i))
	}

	assert.Equal(1001, len(bc.blocks), "Blockchain should handle a large number of blocks")
	assert.Equal("Block 999", string(bc.blocks[len(bc.blocks)-1].Data), "Last block should have correct data")
}
