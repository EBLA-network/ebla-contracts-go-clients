package main

import (
	"log"

	"github.com/EBLA-network/ebla-contracts-go-clients/clients/client_base"
	ebla_client "github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla"
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/ebla_net_config"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	log.Print("Ebla client demo")

	netConfig, err := ebla_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	var config ebla_client.EblaClientConfig
	config.NetConfig = *netConfig
	config.DposContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")

	eblaClient, err := ebla_client.NewEblaClient(config, client_base.WebSocket)
	if err != nil {
		log.Fatal(err)
	}

	totalEligibleVotesCount, err := eblaClient.DposContractClient.GetTotalEligibleVotesCount()
	if err != nil {
		log.Print("GetTotalEligibleVotesCount err: ", err)
	} else {
		log.Printf("GetTotalEligibleVotesCount: %d\n\n", totalEligibleVotesCount)
	}
}
