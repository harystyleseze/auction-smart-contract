package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	NodeURL              string
	ChainID              string
	AuctionContractAddr  string
	NFTContractAddr      string
	PrivateKey          string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
		return nil, err
	}

	return &Config{
		NodeURL:              os.Getenv("NODE_URL"),
		ChainID:              os.Getenv("CHAIN_ID"),
		AuctionContractAddr:  os.Getenv("AUCTION_CONTRACT_ADDRESS"),
		NFTContractAddr:      os.Getenv("NFT_CONTRACT_ADDRESS"),
		PrivateKey:          os.Getenv("PRIVATE_KEY"),
	}, nil
}