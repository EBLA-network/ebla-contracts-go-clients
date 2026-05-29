package ebla_net_config

import (
	"errors"
	"math/big"

	"github.com/EBLA-network/ebla-contracts-go-clients/clients/client_base"
	"github.com/ethereum/go-ethereum/common"
)

func GenNetConfig(network client_base.Network) (*client_base.NetConfig, error) {
	config := new(client_base.NetConfig)

	switch network {
	case client_base.Mainnet:
		config.HttpUrl = "https://rpc.mainnet.ebla.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(841)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case client_base.Testnet:
		config.HttpUrl = "https://rpc.testnet.ebla.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(842)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	case client_base.Devnet:
		config.HttpUrl = "https://rpc.devnet.ebla.io"
		config.WsUrl = "" // TODO
		config.ChainID = big.NewInt(843)
		config.ContractAddress = common.HexToAddress("0x00000000000000000000000000000000000000FE")
		break
	default:
		return nil, errors.New("Invalid network argument")
	}

	return config, nil
}
