package Controller

import "Auction/Blockchain"

/*
	Controller for handle all routes
*/
type Controller struct {
	blockChain     *Blockchain.BlockChain
	currentNodeUrl string
}
