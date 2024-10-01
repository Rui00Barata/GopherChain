package block

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Unit test for NewBlock function
func TestNewBlock(t *testing.T) {
	assert := assert.New(t)

	data := "Test Block"
	prevBlockHash := []byte("previous hash")

	block := NewBlock(data, prevBlockHash)

	assert.NotNil(block, "NewBlock should not return nil")
	assert.Equal(data, string(block.Data), "Block data should match input")
	assert.Equal(prevBlockHash, block.PrevBlockHash, "Previous block hash should match input")
	assert.NotEmpty(block.Hash, "Block hash should not be empty")
}

// Unit test for SetHash function
func TestSetHash(t *testing.T) {
	assert := assert.New(t)

	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte("Test Data"),
		PrevBlockHash: []byte("previous hash"),
		Hash:          []byte{},
	}

	block.SetHash()

	assert.NotEmpty(block.Hash, "SetHash should set a non-empty hash")
}

// Integration test for block creation and hash setting
func TestBlockIntegration(t *testing.T) {
	assert := assert.New(t)

	data := "Integration Test Block"
	prevBlockHash := []byte("previous hash")

	block := NewBlock(data, prevBlockHash)

	assert.Equal(data, string(block.Data), "Block data should match input")
	assert.Equal(prevBlockHash, block.PrevBlockHash, "Previous block hash should match input")
	assert.NotEmpty(block.Hash, "Block hash should not be empty")

	// Verify timestamp is recent
	now := time.Now().UnixNano()
	assert.True(block.Timestamp <= now && block.Timestamp > now-(60*1e9), "Block timestamp should be within the last minute")

	// Verify hash changes when data changes
	originalHash := make([]byte, len(block.Hash))
	copy(originalHash, block.Hash)

	block.Data = []byte("Modified Data")
	block.SetHash()

	assert.NotEqual(originalHash, block.Hash, "Block hash should change when data is modified")

	// Test hash consistency
	data = "Consistency Test Block"
	prevBlockHash = []byte("previous hash")

	block1 := NewBlock(data, prevBlockHash)
	block2 := NewBlock(data, prevBlockHash)

	assert.NotEqual(block1.Hash, block2.Hash, "Blocks with same data but created at different times should have different hashes")
}
