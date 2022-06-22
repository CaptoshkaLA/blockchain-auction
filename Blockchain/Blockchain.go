package Blockchain

import (
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"strconv"
	"time"
)

func (b *BlockChain) RegisterBid(bid Bid) {
	b.PendingBids = append(b.PendingBids, bid)
}

func (b *BlockChain) RegisterNode(node string) bool {
	// Add node if it does not exist, else do nothing
	if !b.NetworkNodes[node] {
		b.NetworkNodes[node] = true
		return true // node added
	}
	return false // node already exists
}

func (b *BlockChain) GetLastBlock() Block {
	return b.ChainOfBlocks[len(b.ChainOfBlocks)-1]
}

/*
	CreateNewBlock is a method for creating a new block and insert it
*/
func (b *BlockChain) CreateNewBlock(nonce int, previousBlockHash string, hash string) Block {
	newBlock := Block{
		BlockId:           len(b.ChainOfBlocks) + 1, // Length of chain + 1
		BlockTimestamp:    time.Now().UnixNano(),
		Bids:              b.PendingBids,
		Nonce:             nonce,
		Hash:              hash,
		PreviousBlockHash: previousBlockHash,
	}

	b.PendingBids = Bids{}

	b.ChainOfBlocks = append(b.ChainOfBlocks, newBlock)

	return newBlock
}

/*
	HashBlock calculates hash for block
*/
func (b *BlockChain) HashBlock(previousBlockHash string, currentBlockData string, nonce int) string {

	var stringToHash string = previousBlockHash + currentBlockData + strconv.Itoa(nonce)

	var hash hash.Hash = sha256.New()
	hash.Write([]byte(stringToHash))

	var base64Hash string = base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return base64Hash
}

/*
	PoW is consensus protocol underlying the blockchain
*/
func (b *BlockChain) ProofOfWork(previousBlockHash string, currentBlockData string) int {

	nonce := -1
	inputFormat := ""

	// String starting with “0000”
	for inputFormat != "0000" {
		nonce = nonce + 1
		var hashed string = b.HashBlock(previousBlockHash, currentBlockData, nonce)
		inputFormat = hashed[0:4]
	}

	return nonce
}

/*
	CheckNewBlockHash is a method for validating new block
*/
func (b *BlockChain) CheckNewBlockHash(newBlock Block) bool {
	var lastBlock Block = b.GetLastBlock()
	return lastBlock.Hash == newBlock.PreviousBlockHash &&
		lastBlock.BlockId == newBlock.BlockId-1
}
