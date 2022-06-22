package Blockchain

/*
	BlockChain is a classic implementation includes chain of blocks, bids and nodes in network
*/
type BlockChain struct {
	ChainOfBlocks Blocks          `json:"chain"`
	PendingBids   Bids            `json:"pending_bids"`
	NetworkNodes  map[string]bool `json:"network_nodes"`
}
