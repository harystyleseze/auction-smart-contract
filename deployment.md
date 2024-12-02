harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ forge build
[⠢] Compiling...
No files changed, compilation skipped

harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ forge create --rpc-url localhost:8545 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 src/MyNFT.sol:MyNFT --constructor-args 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
[⠒] Compiling...
No files changed, compilation skipped
Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Deployed to: 0x5FbDB2315678afecb367f032d93F642f64180aa3
Transaction hash: 0xb8a4f96df9f7df6de1a02c14d2c08b9a43a0ae452f1b7601b1d81b4874f2b163
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$

=========================================


harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "safeMint(address, uint256)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 1

blockHash               0xcd423df538ad631dfde6a7daafa106aaf1e8829845457a624f352ad852bd6bc7
blockNumber             2
contractAddress         
cumulativeGasUsed       71240
effectiveGasPrice       883660910
from                    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
gasUsed                 71240
logs                    [{"address":"0x5fbdb2315678afecb367f032d93f642f64180aa3","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000000000000000000000000000000000000000000000","0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266","0x0000000000000000000000000000000000000000000000000000000000000001"],"data":"0x","blockHash":"0xcd423df538ad631dfde6a7daafa106aaf1e8829845457a624f352ad852bd6bc7","blockNumber":"0x2","blockTimestamp":"0x674ca7be","transactionHash":"0x390093ce8f6506bbcbb55869356df14cddd6293e69271beb2287b58718449d0f","transactionIndex":"0x0","logIndex":"0x0","removed":false}]
logsBloom               0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000008000000000000000000040000000000000000000000000040020000000000000100000800000000000000000000000010000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000000000000000042000000200000000000000000000000002000000000000000000060000000000000000000000000000000000000000000000000000000000000000000
root                    
status                  1 (success)
transactionHash         0x390093ce8f6506bbcbb55869356df14cddd6293e69271beb2287b58718449d0f
transactionIndex        0
type                    2
blobGasPrice            1
blobGasUsed             
authorizationList       
to                      0x5FbDB2315678afecb367f032d93F642f64180aa3
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 


============================================

harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ forge create --rpc-url localhost:8545 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 src/EnglishAuction.sol:EnglishAuction --constructor-args 0x5fbdb2315678afecb367f032d93f642f64180aa3 2
[⠒] Compiling...
No files changed, compilation skipped
Deployer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Deployed to: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
Transaction hash: 0x1e1d243c37dd98ff3f12e2e450ed2150d9066320bbfc5ee8244e1988ac75b528
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 


-------------------------
Approve the english auction contract to make changes on our behalf as the owner after minting the NFT

cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "approve(address, uint256)" 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0 2


harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ cast send 0x5FbDB2315678afecb367f032d93F642f64180aa3 --unlocked --from 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "approve(address,
 uint256)" 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0 2

blockHash               0x56805f17b0d338bc8f31c36fdbf65f8f675f533c1ecf67cc450f06edf3664cb4
blockNumber             4
contractAddress         
cumulativeGasUsed       48685
effectiveGasPrice       680877939
from                    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
gasUsed                 48685
logs                    [{"address":"0x5fbdb2315678afecb367f032d93f642f64180aa3","topics":["0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925","0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266","0x0000000000000000000000009fe46736679d2d9a65f0992f2272de9f3c7fa6e0","0x0000000000000000000000000000000000000000000000000000000000000002"],"data":"0x","blockHash":"0x56805f17b0d338bc8f31c36fdbf65f8f675f533c1ecf67cc450f06edf3664cb4","blockNumber":"0x4","blockTimestamp":"0x674d6df5","transactionHash":"0x13cfbc958704e2973314b6f281bf39ca1c9debc4cb157991bc9dac99397789a7","transactionIndex":"0x0","logIndex":"0x0","removed":false}]
logsBloom               0x04000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000000400000000000000000000000200000000000000000000000000000000000000000000000000000000000000000400000000040000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000020000000000000000000100000000000000000000000000000000000000000000000040000000200000000000000000000000002000000000000000000000000010000000000000000000000000000000000000008000000000000000000000
root                    
status                  1 (success)
transactionHash         0x13cfbc958704e2973314b6f281bf39ca1c9debc4cb157991bc9dac99397789a7
transactionIndex        0
type                    2
blobGasPrice            1
blobGasUsed             
authorizationList       
to                      0x5FbDB2315678afecb367f032d93F642f64180aa3
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 

=============================

Start API
go run . in api folder

curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"opening_bid": 1, "duration": 1200}' localhost:8080/auction/start


curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"opening_bid": 1, "duration": 1200}' localhost:8080/auction/start
0xd315dca76d22d1d4fb8bd0b048f6f8495d16b00462ec60914407e3b85211516charystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 

Start auction
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"opening_bid": 1, "duration": 1200}' localhost:8080/auction/start
execution reverted: revert: Auction has already startedharystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 

place bid

curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"value": 1}' localhost:8080/bids

harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"opening_bid": 1, "duration": 1200}' localhost:8080/auction/start
execution reverted: revert: Auction has already startedharystcurl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"value": 1}' localhost:8080/bids
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"value": 3}' localhost:8080/bids
0xe16e27f8698f90ad65ca6b7f6911e19e43c0c029202a72061f1e83660175816aharystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 

Place bid as someone else
curl -H "Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" -X POST --data '{"value": 3}' localhost:8080/bids


harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST --data '{"value": 3}' localhost:8080/bids
0xe16e27f8698f90ad65ca6b7f6911e19e43c0c029202a72061f1e8366017curl -H "Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" -X POST --data '{"value": 3}' localhost:8080/bids
harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" -X POST --data '{"value": 20}' localhost:8080/bid
s
0x50ffb0621cd7d2e925c35838c905d6375e98c6ffbd145b9db1a51b033c5c83a0harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ 

Withdraw bid
curl -H "Key: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST localhost:8080/auction/withdraw

harystyles@DESKTOP-VM8UCIK:~/smart_contract/english-auction$ curl -H "Key: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST localhost:8080/auction/withdraw
0xcd12cb95ab2036373463b9af69ce6662de3f21833c1e3590c816e1ccc071bd6a


End bid as the owner
curl -H "Key: ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" -X POST localhost:8080/auction/end

End bid as external person
curl -H "Key: 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" -X POST localhost:8080/auction/end

0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

-------------------------

Deploy to testnet

get the rpc url

forge create --rpc-url https://eth-sepolia.g.alchemy.com --private-key $PRIVATE_KEY src/MyNFT.sol:MyNFT --constructor-args yourAddress(owner of the nft)

Mint the nft

Deploy Auction contract 

forge create --rpc-url https://eth-sepolia.g.alchemy.com --private-key $PRIVATE_KEY src/EnglishAuction.sol:EnglishAuction --constructor-args <deployed to contract address> <nftId>


# 1. First approve the NFT transfer
curl -H "Key: YOUR_PRIVATE_KEY" -X POST localhost:8080/nft/approve

# 2. Then start the auction
curl -H "Key: YOUR_PRIVATE_KEY" -X POST --data '{"opening_bid": 2, "duration": 1200}' localhost:8080/auction/start

===================

godotenv package
go get github.com/joho/godotenv