package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	contract "api/auction"
)

func main() {
	// Load the configuration
	config, err := LoadConfig()
	// Return the error if the configuration is invalid
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
		return
	}

	// Create a new Gin router
	r := gin.Default()
	// Connect to the Ethereum node
	conn, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		log.Err(err).Msg("Fail to initialise the eth client")
		return
	}

	// Load the smart contract
	auction, err := contract.NewAuction(common.HexToAddress(config.AuctionContractAddr), conn)
	if err != nil {
		log.Err(err).Msg("Fail to initialise the contract")
		return
	}

	// Initialise the handler
h := NewHandler(auction, config, conn)
	// Register the routes
	r.POST("/nft/approve", h.approveNFT)
	r.POST("/auction/start", h.startAuction)
	r.POST("/auction/end", h.endAuction)
	r.POST("/auction/withdraw", h.withdraw)
	r.POST("/bids", h.createBid)

	// Start the server
	if err := r.Run(); err != nil {
		log.Err(err).Msg("Server exited")
	}
}
