package Model

/*
	Bid implements information about bid
*/
type Bid struct {
	OwnerName string  `json:"bidder_name"`
	AuctionId int     `json:"auction_id"`
	BidValue  float32 `json:"bid_value,string"`
}
type Bids []Bid
