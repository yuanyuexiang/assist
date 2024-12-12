package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/log"
)

func main() {
	//https://linea-sepolia.infura.io/v3/<YOUR-API-KEY>
	//wss://linea-sepolia.infura.io/ws/v3/<YOUR-API-KEY>
	//url:= "https://linea-sepolia.infura.io/v3/dccc894f6609403fa5e005bd1b228f88"
	url2 := "wss://linea-sepolia.infura.io/ws/v3/dccc894f6609403fa5e005bd1b228f88"
	//conn, err := ethclient.Dial("https://mantle-sepolia.infura.io/v3/dccc894f6609403fa5e005bd1b228f88")
	conn, err := ethclient.Dial(url2)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	// NOTE update the deployment address!

	address := "0x5e817cd749adcbc8a0c6d54ef0efcfe34346bce7"
	//address := common.HexToAddress("0x7e3dC1d6DA15692340e784D803232993e8a3366F")
	store, err := NewStorage(common.HexToAddress(string(address)), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	status, err := store.GetStatus(nil, common.HexToAddress("0xadA778c33B4CA3f5374D410396b84DE2B08CC567"))
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
	log.Printf("Status: %s", status)

	eventCh := make(chan *StorageStatusUpdated)
	opts := &bind.WatchOpts{}
	go func() {
		for event := range eventCh {
			log.Printf("Status updated: %s", event.NewStatus)
		}
	}()
	_, err = store.WatchStatusUpdated(opts, eventCh, nil)
	if err != nil {
		log.Fatalf("Failed to watch status updated: %v", err)
	}
	log.Println("Listening for events...")
	select {}
}
