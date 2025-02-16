package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int    // Position of the block in the chain
	Timestamp    string // Time when the block was created
	Data         string // Transaction data
	PrevHash     string // Hash of the previous block
	Hash         string // Hash of this block
	Nonce        int    // Proof of Work variable
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

const difficulty = 4 // Adjust difficulty by changing the number of leading zeros

// calculateHash generates a SHA-256 hash for a block
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// generateBlock mines a new block with Proof of Work
func generateBlock(oldBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Nonce = 0

	// Proof of Work: Find a hash with `difficulty` leading zeros
	for {
		newBlock.Hash = calculateHash(newBlock)
		if strings.HasPrefix(newBlock.Hash, strings.Repeat("0", difficulty)) {
			break
		}
		newBlock.Nonce++
	}
	
	return newBlock
}

// isBlockValid ensures the new block is properly linked to the previous block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// addBlock appends a valid new block to the blockchain
func addBlock(newBlock Block) {
	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
	}
}

// createGenesisBlock initializes the blockchain with the first block
func createGenesisBlock() Block {
	genesisBlock := Block{Index: 0, Timestamp: time.Now().String(), Data: "Genesis Block", PrevHash: "", Nonce: 0}
	genesisBlock.Hash = calculateHash(genesisBlock)
	return genesisBlock
}

func main() {
	// Initialize blockchain with the genesis block
	Blockchain = append(Blockchain, createGenesisBlock())

	// Add new blocks with some data
	newBlock1 := generateBlock(Blockchain[len(Blockchain)-1], "Alice sent 1 BTC to Bob")
	addBlock(newBlock1)

	newBlock2 := generateBlock(Blockchain[len(Blockchain)-1], "Bob sent 0.5 BTC to Charlie")
	addBlock(newBlock2)

	// Print the blockchain
	for _, block := range Blockchain {
		fmt.Printf("\nIndex: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\nNonce: %d\n",
			block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Nonce)
	}
}
