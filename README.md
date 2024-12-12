
接入点
https://developer.metamask.io/key/dccc894f6609403fa5e005bd1b228f88/active-endpoints

合约绑定
使用go-ethereum
https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings

$ solc --abi Storage.sol -o build

$ abigen --abi build/Storage.abi --pkg main --type Storage --out Storage.go
