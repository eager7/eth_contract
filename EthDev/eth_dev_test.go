package eth_dev_test

import (
	"fmt"
	"github.com/eager7/eth_contract/EthDev"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	"testing"
)

func TestAbi(t *testing.T) {
	devAbi, err := abi.JSON(strings.NewReader(string(eth_dev.EthDevABI)))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(devAbi)
}
