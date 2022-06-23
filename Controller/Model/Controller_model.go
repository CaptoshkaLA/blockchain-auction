package Model

import "Auction/Blockchain"

/*
	Controller for handle all routes
*/
type Controller struct {
	BlockChain     *Blockchain.BlockChain
	CurrentNodeUrl string
}
