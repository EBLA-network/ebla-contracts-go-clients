package main

import (
	"log"

	"github.com/EBLA-network/ebla-contracts-go-clients/clients/client_base"
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/rpc_client"
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/ebla_net_config"
)

func main() {
	log.Print("Ebla rpc client demo")

	config, err := ebla_net_config.GenNetConfig(client_base.Testnet)
	if err != nil {
		log.Fatal(err)
	}

	// var config client_base.NetConfig
	// config.HttpUrl = "http://localhost:7017"
	// config.ChainID = big.NewInt(649)

	rpcClient, err := rpc_client.NewRpcClient(*config, client_base.Http)
	if err != nil {
		log.Fatal(err)
	}

	pillarBlockData, err := rpcClient.GetPillarBlockData(100, true)
	if err != nil {
		log.Fatal("GetPillarBlockData err: ", err)
	} else {
		log.Printf("GetPillarBlockData: %d\n\n", pillarBlockData)
		log.Printf("pillarBlockData.PillarBlock.PbftPeriod: %d\n\n", pillarBlockData.PillarBlock.PbftPeriod)
	}

	eblaConfig, err := rpcClient.GetEblaConfig()
	if err != nil {
		log.Fatal("GetEblaConfig err: ", err)
	} else {
		log.Printf("GetEblaConfig: %d\n\n", eblaConfig)
		log.Printf("eblaConfig.Hardforks.FicusHf.PillarBlocksInterval: %d\n\n", uint64(eblaConfig.Hardforks.FicusHf.PillarBlocksInterval))
	}
}
