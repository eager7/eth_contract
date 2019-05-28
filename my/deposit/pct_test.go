package deposit

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
)

func TestDeposit(t *testing.T) {
	client, err := ethclient.Dial("http://127.0.0.1:7545") //"wss://mainnet.infura.io/ws"
	if err != nil {
		t.Fatal(err)
	}
	contract := common.HexToAddress("0x4cfcd906dfea370d9d7cc62b48b7147ca60f0f99")
	instance, err := NewDeposit(contract, client)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := instance.Balance(&bind.CallOpts{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("the contract balance :", balance)

	cAbi, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("withdrawal:", cAbi.Methods["withdrawal"])
	data, err := cAbi.Pack("withdrawal", common.Address{}, new(big.Int).SetUint64(1))
	if err != nil {
		t.Fatal(err)
	}

	call := ethereum.CallMsg{
		From:     common.Address{},
		To:       &contract,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     data,
	}
	limit, err := client.EstimateGas(context.Background(), call)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("limit:", limit)
}
