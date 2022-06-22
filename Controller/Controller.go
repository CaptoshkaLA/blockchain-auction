package Controller

import (
	. "Auction/Blockchain"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetBlockChain GET represents the blockchain in JSON format. Output be like:
/*   {
	"chain": [
		{
			"index": 1,
			"timestamp": 1111111111111111111,
			"bids": [],
			"nonce": 111,
			"hash": "0",
			"previous_block_hash": "0"
		}
	],
	"pending_bids": [],
	"network_nodes": []
}
*/
func (c *Controller) GetBlockChain(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(c.blockChain)

	writer.Write(data)
	return
}

// Consensus GET /consensus
/* Consensus ensures that this node - and then all the network — have the same chains,
   with the same bets: The network which contains the longest chain keeps it, forcing the
   other to drop its chain and get the new one*/
func (c *Controller) Consensus(writer http.ResponseWriter, request *http.Request) {

	// Iterate over all nodes, getting each node's blockchain and measuring its length
	// to identify the longest chain
	for key, _ := range c.blockChain.NetworkNodes {
		// Ignore this node
		if key == c.currentNodeUrl {
			continue
		}

		// Call /blockchain on the current node
		requestUrl, _ := url.Parse(key + "/blockchain")
		request := &http.Request{
			Method: "GET",
			URL:    requestUrl, // URL *url.URL
			Header: http.Header{ // type Header map[string][]string
				"Content-Type": {"application/json"},
			},
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Printf("Failed to call /blocchain on node %s. Error: %s", key, err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		// Process response from node which is the node's blockchain
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Failed to call /blocchain on node %s. Error: %s", key, err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		var blockChain *BlockChain
		err = json.Unmarshal(body, blockChain)
		if err != nil {
			log.Printf("Failed to process response from  node %s. Error: %s", key, err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Get length of this chain, and update maximum length if necessary

	}
}

/* Helpers */
func (c *Controller) broadcastToAllNodes(api string, body []byte) {
	for key, _ := range c.blockChain.NetworkNodes {
		if key != c.currentNodeUrl {
			doPostCall(key+api, body)
		}
	}
}

// Do a post call to the given url. Typically used to inform other nodes of interesting changes
// such as a new block or a new node
func doPostCall(url string, body []byte) error {
	contentType := "application/json;charset=UTF-8"

	var buffer *bytes.Buffer = bytes.NewBuffer(body)
	response, err := http.Post(url, contentType, buffer)
	if err != nil {
		log.Printf("Failed to POST call to %s: %s", url, err)
		return err
	}

	response.Body.Close()
	return nil
}
