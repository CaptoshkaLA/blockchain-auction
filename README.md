
> # About
 <hr>
 This work is devoted to the creation of an online auction based on the blockchain network. <br/> The consensus underlying the network is PoW. <br/>
 <hr/>

> # Setup
 <hr>

### `Go`
To install Go visit site:

[`https://go.dev/doc/install`](https://go.dev/doc/install) <br />

### `Gorilla`
To install gorilla/mux:

`go get github.com/gorilla/mux` <br />

 <hr/>
 
 
> # Running
> <hr>
 `go run Main.go <port_to_listen>` (For example `go run Main.go 9000`) <br />
 
 Run postman and invoke API Methods
 <hr/>
 
> # Postman
 <hr>
 You can use Postman to work with the API.

### `Install`
To install Postman:

`$ sudo apt update` <br />
`$ sudo apt install snapd` <br />
`$ sudo snap install postman` <br />

### `Making a POST call`

Path: <br/> 
`localhost:9000/bid/broadcast` <br/>

POST body: <br/> 
`{` <br/> 
 ` "bidder_name": "Andrew",` <br/> 
 ` "auction_id": 1,` <br/> 
 ` "bid_value,string": "1.25"` <br/> 
`}` <br/> 

Response body: <br/> 
`{`  <br/> 
`"Name": "RegisterAndBroadcastBid"`, <br/> 
`"Status": "Bid created and broadcast successfully",` <br/> 
`"Time": "2022-06-23T19:12:08.88213331+03:00"` <br/> 
`}`<br/> 

### `Making a GET call`

To check, you can use the link and see the pending bids.
Path: <br/> 
`localhost:9000/blockchain` <br/>

Response body: <br/> 
`{` <br/> 
`"chain": [` <br/> 
`{` <br/> 
`"block_id": 1,` <br/> 
`"block_timestamp": 1656001600038743150,` <br/> 
`"bids": [],` <br/> 
`"nonce": 100,` <br/> 
`"hash": "0",` <br/> 
`"previous_block_hash": "0"` <br/> 
`}` <br/> 
`],` <br/>  
`"pending_bids": [` <br/> 
`{` <br/> 
`"bidder_name": "Anton",` <br/> 
`"auction_id": 1,` <br/> 
`"bid_value": "1.25"` <br/> 
`}` <br/> 
`],` <br/> 
`"network_nodes": {}` <br/> 
`}` <br/> 

 <hr/>


