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

This project is licensed under the MIT License - see the LICENSE file for details.
```

This README provides a comprehensive overview of your project while maintaining a clean and professional format. It includes all the essential sections: overview, installation, usage instructions, development commands, and contribution guidelines.

The structure is based on common patterns seen in the codebase, particularly referencing:


````1:66:english-auction/README.md
## Foundry

**Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust.**

Foundry consists of:

-   **Forge**: Ethereum testing framework (like Truffle, Hardhat and DappTools).
-   **Cast**: Swiss army knife for interacting with EVM smart contracts, sending transactions and getting chain data.
-   **Anvil**: Local Ethereum node, akin to Ganache, Hardhat Network.
-   **Chisel**: Fast, utilitarian, and verbose solidity REPL.

## Documentation

https://book.getfoundry.sh/

## Usage

### Build

```shell
$ forge build
```

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
````


And incorporating deployment information from:


```58:59:deployment.md
# 2. Then start the auction
curl -H "Key: YOUR_PRIVATE_KEY" -X POST --data '{"opening_bid": 2, "duration": 1200}' localhost:8080/auction/start
