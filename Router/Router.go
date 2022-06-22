package Router

import (
	. "Auction/Blockchain"
	. "Auction/Controller"
)

/*
	Instantiate a controller object so that routes can be initialized
*/
var controller *Controller = &Controller{
	BlockChain: &BlockChain{
		ChainOfBlocks: Blocks{},
		PendingBids:   Bids{},
		NetworkNodes:  map[string]bool{},
	},
	CurrentNodeUrl: "",
}

/*
	Define all routes (name, http method, path, and controller api)
*/
var routes []Route = []Route{
	Route{
		Name:        "Index",
		Method:      "GET",
		Path:        "/",
		HandlerFunc: controller.Index,
	},
	Route{
		Name:        "GetBlockChain",
		Method:      "GET",
		Path:        "/blockchain",
		HandlerFunc: controller.GetBlockChain,
	},
	Route{
		Name:        "RegisterAndBroadcastBid",
		Method:      "POST",
		Path:        "/bid/broadcast",
		HandlerFunc: controller.RegisterAndBroadcastBid,
	},
	Route{
		Name:        "RegisterBid",
		Method:      "POST",
		Path:        "/bid",
		HandlerFunc: controller.RegisterBid,
	},
	Route{
		Name:        "RegisterAndBroadcastNode",
		Method:      "POST",
		Path:        "/register-and-broadcast-node",
		HandlerFunc: controller.RegisterAndBroadcastNode,
	},
	Route{
		Name:        "RegisterNode",
		Method:      "POST",
		Path:        "/register-node",
		HandlerFunc: controller.RegisterNode,
	},
	Route{
		Name:        " RegisterNodesBulk",
		Method:      "POST",
		Path:        "/register-nodes-bulk",
		HandlerFunc: controller.RegisterNodesBulk,
	},
	Route{
		Name:        "Mine",
		Method:      "GET",
		Path:        "/mine",
		HandlerFunc: controller.Mine,
	},
	Route{
		Name:        "ReceiveNewBlock",
		Method:      "POST",
		Path:        "/receive-new-block",
		HandlerFunc: controller.ReceiveNewBlock,
	},
	Route{
		Name:        "Consensus",
		Method:      "GET",
		Path:        "/consensus",
		HandlerFunc: controller.Consensus,
	},
	Route{
		Name:        "GetBidsForAuction",
		Method:      "GET",
		Path:        "/auction/{auctionId}",
		HandlerFunc: controller.GetBidsForAuction,
	},
	Route{
		Name:        "GetBidsForPlayer",
		Method:      "GET",
		Path:        "/player/{playerId}",
		HandlerFunc: controller.GetBidsForPlayer,
	},
}
