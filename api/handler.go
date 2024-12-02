// Handles the actual business logic for API endpoints
package main

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	contract "api/auction"
)

// Create a handler
type Handler struct {
	auction *contract.Auction
	config  *Config
	client  bind.ContractBackend
}

// Create a new handler
func NewHandler(auction *contract.Auction, config *Config, client bind.ContractBackend) *Handler {
	return &Handler{
		auction: auction,
		config:  config,
		client:  client,
	}
}

// Create transaction options
func (h *Handler) txOpts(c *gin.Context) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(h.config.PrivateKey)
	if err != nil {
		return nil, err
	}

	chainID, ok := new(big.Int).SetString(h.config.ChainID, 10)
	if !ok {
		return nil, errors.New("invalid chain ID")
	}

	return bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}

// Start the auction
func (h *Handler) startAuction(c *gin.Context) {
	// Create transaction options
	opts, err := h.txOpts(c)
	// Return the error to the client if the transaction options are invalid
	if err != nil {
		log.Err(err).Msg("Failed to create opts")
		c.String(http.StatusBadRequest, "invalid auth")
		return
	}

	// Parse the request
	var req StartAuctionRequest
	// Return the error to the client if the request is invalid
	if err := c.ShouldBindJSON(&req); err != nil {
		// Log the error
		log.Err(err).Msg("invalid request body")
		// Return the error to the client
		c.String(http.StatusBadRequest, "invalid request body")
		return
	}

	// Call the smart contract to start the auction
	tx, err := h.auction.Start(opts, big.NewInt(req.OpeningBid), big.NewInt(req.Duration))
	// Return the error to the client if the auction fails
	if err != nil {
		log.Err(err).Msg("Error occurred while starting the auction")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, tx.Hash().String())
}

// End the auction
func (h *Handler) endAuction(c *gin.Context) {
	opts, err := h.txOpts(c)
	// Return the error to the client if the transaction options are invalid
	if err != nil {
		log.Err(err).Msg("Failed to create txOpts")
		c.String(http.StatusBadRequest, "invalid auth")
		return
	}

	// Call the smart contract to end the auction
	tx, err := h.auction.End(opts)
	if err != nil {
		log.Err(err).Msg("Error occurred while starting the auction")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, tx.Hash().String())
}

func (h *Handler) withdraw(c *gin.Context) {
	// Create transaction options
	opts, err := h.txOpts(c)
	// Return the error to the client if the transaction options are invalid
	if err != nil {
		log.Err(err).Msg("Failed to create txOpts")
		c.String(http.StatusBadRequest, "invalid auth")
		return
	}
	// Call the smart contract to withdraw
	tx, err := h.auction.Withdraw(opts)
	// Return the error to the client if the withdrawal fails
	if err != nil {
		log.Err(err).Msg("Error occurred while starting the auction")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, tx.Hash().String())
}

// Create a bid on the auction
func (h *Handler) createBid(c *gin.Context) {
	// Parse the request
	var req BidRequest
	// Return the error to the client if the request is invalid
	if err := c.ShouldBindJSON(&req); err != nil {
		// Log the error
		log.Err(err).Msg("invalid request body")
		// Return the error to the client
		c.String(http.StatusBadRequest, "invalid request body")
		return
	}

	// Create transaction options
	opts, err := h.txOpts(c)
	// Return the error to the client if the transaction options are invalid
	if err != nil {
		// Log the error
		log.Err(err).Msg("Failed to create txOpts")
		// Return the error to the client
		c.String(http.StatusBadRequest, "invalid auth")
		return
	}

	// Set the value of the transaction to the bid amount
	opts.Value = big.NewInt(req.Value)

	// Call the smart contract to create a bid
	tx, err := h.auction.Bid(opts)
	if err != nil {
		// Log the error
		log.Err(err).Msg("Error occurred while starting the auction")
		// Return the error to the client
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Return the transaction hash to the client
	c.String(http.StatusOK, tx.Hash().String())
}

func (h *Handler) approveNFT(c *gin.Context) {
	opts, err := h.txOpts(c)
	if err != nil {
		log.Err(err).Msg("Failed to create opts")
		c.String(http.StatusBadRequest, "invalid auth")
		return
	}

	// Create NFT contract instance
	nft, err := contract.NewMyNFT(common.HexToAddress(h.config.NFTContractAddr), h.client)
	if err != nil {
		log.Err(err).Msg("Failed to create NFT contract instance")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Approve auction contract to transfer NFT
	tx, err := nft.Approve(opts, common.HexToAddress(h.config.AuctionContractAddr), big.NewInt(1))
	if err != nil {
		log.Err(err).Msg("Failed to approve NFT transfer")
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, tx.Hash().String())
}
