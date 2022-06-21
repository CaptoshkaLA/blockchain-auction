package Model

/*
	Block is a classic part of blockchain
*/
type Block struct {
	BlockId           int    `json:"block_id"`
	BlockTimestamp    int64  `json:"block_timestamp"`
	Bids              Bids   `json:"bids"`
	Nonce             int    `json:"nonce"`
	Hash              string `json:"hash"`
	PreviousBlockHash string `json:"previous_block_hash"`
}
type Blocks []Block

/*
	Auxiliary struct required in mining
*/
type BlockData struct {
	Index string
	Bids  Bids
}
