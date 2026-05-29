package ebla_client

import (
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/client_base"
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/dpos_contract_client"
	"github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/rpc_client"
	"github.com/ethereum/go-ethereum/common"
)

type EblaClient struct {
	*client_base.ClientBase
	DposContractClient *dpos_contract_client.DposContractClient
	RpcClient          *rpc_client.RpcClient
}

type EblaClientConfig struct {
	client_base.NetConfig
	DposContractAddress common.Address `json:"dpos_contract_address"`
}

func NewEblaClient(config EblaClientConfig, communicationProtocol client_base.CommunicationProtocol) (*EblaClient, error) {
	var err error

	eblaClient := new(EblaClient)
	eblaClient.ClientBase, err = client_base.NewClientBase(config.NetConfig, communicationProtocol)
	if err != nil {
		return nil, err
	}

	eblaClient.DposContractClient, err = dpos_contract_client.NewSharedDposContractClient(eblaClient.ClientBase, config.DposContractAddress)
	if err != nil {
		return nil, err
	}

	eblaClient.RpcClient = rpc_client.NewSharedRpcClient(eblaClient.ClientBase)

	return eblaClient, nil
}
