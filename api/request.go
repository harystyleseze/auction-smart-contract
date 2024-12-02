// Defines the structure of incoming API requests
package main

// Create a start auction request
type StartAuctionRequest struct {
	OpeningBid int64 `json:"opening_bid"`
	Duration   int64 `json:"duration"`
}

// Create a bid request
type BidRequest struct {
	Value int64 `json:"value"`
}
