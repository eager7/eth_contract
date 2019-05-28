package pct

import (
	"context"
	"encoding/hex"
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
	instance, err := NewPct(contract, client)
	if err != nil {
		t.Fatal(err)
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, common.Address{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("the contract balance :", balance)

	cAbi, err := abi.JSON(strings.NewReader(PctABI))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("transfer:", cAbi.Methods["transfer"])
	data, err := cAbi.Pack("transfer", common.Address{}, new(big.Int).SetUint64(1))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("data hex:", hex.EncodeToString(data))

	call := ethereum.CallMsg{
		From:     common.HexToAddress("0x4cfcd906dfea370d9d7cc62b48b7147ca60f0f99"),
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
