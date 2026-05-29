# ebla-contracts-go-clients

Go client and generated bindings for EBLA's on-chain contracts. Currently it
covers the **DPoS** precompile (`clients/ebla/dpos_contract_client`) plus a
generic RPC client and net config. It builds against upstream
`github.com/ethereum/go-ethereum` (no fork required). EBLA has no Ethereum
bridge, so no bridge/eth clients are included.

# Usage

See `clients/ebla/demo/demo.go` for an example.

#### Prerequisites
##### solc
```
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

##### abigen
```
go install github.com/ethereum/go-ethereum@latest
cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum@v<X.Y.Z>/
make
make devtools
```

## Regenerating the DPoS interface from EBLA contracts

The DPoS contract source lives in the `ebla-evm` repo. Add it as a submodule
under `submodules/ebla-evm` (or point the paths below at a local checkout),
then regenerate the ABI + Go bindings:

```
solc --abi --overwrite --optimize submodules/ebla-evm/ebla/state/contracts/dpos/solidity/dpos_contract_interface.sol --output-dir clients/ebla/dpos_contract_client/dpos_interface/

abigen --abi=clients/ebla/dpos_contract_client/dpos_interface/DposInterface.abi --pkg=dpos_interface --out=clients/ebla/dpos_contract_client/dpos_interface/dpos_interface.go
```
