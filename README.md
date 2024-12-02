# English Auction Smart Contract

A decentralized English auction implementation built with Solidity and Go, using Foundry for smart contract development and testing.

## Overview

This project implements an English auction system where:
- Sellers can auction NFTs
- Bidders can place bids during the auction period
- Highest bidder wins the auction after the duration ends
- Losers can withdraw their bids

## Prerequisites

- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [Node.js](https://nodejs.org/) (for development tools)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd english-auction
```

2. Install dependencies:
```bash
forge install
go mod tidy
```

## Smart Contract Development

### Build
```bash
forge build
```

### Test
```bash
forge test
```

### Format
```bash
forge fmt
```

### Generate Gas Report
```bash
forge snapshot
```

## Go Backend Development

### Generate Contract Bindings

1. Build the contract:
```bash
cd english-auction
forge build
cd ..
```

2. Generate Go bindings:
```bash
jq -r '.abi' english-auction/out/EnglishAuction.sol/EnglishAuction.json > EnglishAuction.abi
jq -r '.bytecode.object' english-auction/out/EnglishAuction.sol/EnglishAuction.json > EnglishAuction.bin
abigen --abi EnglishAuction.abi --bin EnglishAuction.bin --pkg auction --out api/auction/auction.go
```

### Run Backend Server
```bash
go run main.go
```

## API Endpoints

### Start Auction
```bash
curl -H "Key: YOUR_PRIVATE_KEY" -X POST --data '{"opening_bid": <amount>, "duration": <seconds>}' localhost:8080/auction/start
```

## Project Structure

- `/contracts`: Solidity smart contracts
- `/test`: Contract test files
- `/api`: Go backend API code
- `/scripts`: Deployment and utility scripts

## Dependencies

The project uses several key dependencies:
- OpenZeppelin Contracts for secure smart contract components
- Foundry for Ethereum development tools
- Go-Ethereum for blockchain interaction

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License

### Test

```shell
$ forge test
```

### Format

```shell
$ forge fmt
```

### Gas Snapshots

```shell
$ forge snapshot
```

### Anvil

```shell
$ anvil
```

### Deploy

```shell
$ forge script script/Counter.s.sol:CounterScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```

