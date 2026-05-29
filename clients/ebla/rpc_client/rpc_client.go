package rpc_client

import (
	"context"

	"github.com/EBLA-network/ebla-contracts-go-clients/clients/client_base"
	ebla_rpc_types "github.com/EBLA-network/ebla-contracts-go-clients/clients/ebla/rpc_client/types"

	"github.com/ethereum/go-ethereum"
)

type RpcClient struct {
	*client_base.ClientBase
}

func NewRpcClient(config client_base.NetConfig, communicationProtocol client_base.CommunicationProtocol) (*RpcClient, error) {
	var err error

	rpcClient := new(RpcClient)
	rpcClient.ClientBase, err = client_base.NewClientBase(config, communicationProtocol)
	if err != nil {
		return nil, err
	}

	return rpcClient, nil
}

func NewSharedRpcClient(sharedClient *client_base.ClientBase) *RpcClient {
	rpcClient := new(RpcClient)
	rpcClient.ClientBase = sharedClient

	return rpcClient
}

func (RpcClient *RpcClient) GetPillarBlockData(period uint64, includeBinaryData bool) (*ebla_rpc_types.PillarBlockData, error) {
	var pillarBlockData *ebla_rpc_types.PillarBlockData
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &pillarBlockData, "ebla_getPillarBlockData", period, includeBinaryData)
	if err == nil && pillarBlockData == nil {
		err = ethereum.NotFound
	}

	return pillarBlockData, err
}

func (RpcClient *RpcClient) GetEblaConfig() (*ebla_rpc_types.EblaConfig, error) {
	var eblaConfig *ebla_rpc_types.EblaConfig
	err := RpcClient.EthClient.Client().CallContext(context.Background(), &eblaConfig, "ebla_getConfig")
	if err == nil && eblaConfig == nil {
		err = ethereum.NotFound
	}

	return eblaConfig, err
}
