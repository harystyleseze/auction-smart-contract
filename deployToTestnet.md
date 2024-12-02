
forge build

forge create --rpc-url localhost:8545 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 src/MyNFT.sol:MyNFT --constructor-args 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

cast send 0x8A791620dd6260079BF849Dc5567aDC3F2FdC318 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "safeMint(address, uint256)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 1

forge create --rpc-url localhost:8545 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 src/EnglishAuction.sol:EnglishAuction --constructor-args 0x8A791620dd6260079BF849Dc5567aDC3F2FdC318 1

=============================

Deploy to testnet

get the rpc url

# In your terminal or in a script:
set -a
source .env
set +a

# Now run your forge create command
forge create --rpc-url $NODE_URL --private-key $PRIVATE_KEY src/MyNFT.sol:MyNFT --constructor-args 0xbDAA179baF66805f2225f6EA1798C7D889317F6E


[⠢] Compiling...
No files changed, compilation skipped
Deployer: 0xbDAA179baF66805f2225f6EA1798C7D889317F6E
Deployed to: 0x9EA1162288120277fEd5f95bD889C2D3F2a6B396
Transaction hash: 0xf7cb76c947f5451b832d2b05b472196bf7b5bec93e2a08b2a5697257a222e835

Mint the nft
cast send $NFT_CONTRACT_ADDRESS --rpc-url $NODE_URL --private-key $PRIVATE_KEY --from $ADDRESS "safeMint(address,uint256)" $ADDRESS 2

cast send $NFT_CONTRACT_ADDRESS --rpc-url $NODE_URL --private-key $PRIVATE_KEY --from $ADDRESS "safeMint(address,uint256)" 0xbDAA179baF66805f2225f6EA1798C7D889317F6E 2

cast send 0x9EA1162288120277fEd5f95bD889C2D3F2a6B396 --rpc-url $NODE_URL --private-key $PRIVATE_KEY --from $ADDRESS "safeMint(address,uint256)" $ADDRESS 2


Deploy Auction contract 

forge create --rpc-url https://api-holesky.etherscan.io/api --private-key $PRIVATE_KEY src/EnglishAuction.sol:EnglishAuction --constructor-args <deployed to contract address> <nftId>
