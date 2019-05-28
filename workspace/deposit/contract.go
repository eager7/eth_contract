package deposit

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func PackMessage(name string, args ...interface{}) string {
	cAbi, err := abi.JSON(strings.NewReader(string(DepositABI)))
	if err != nil {
		return ""
	}
	data, err := cAbi.Pack(name, args...)
	if err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(data)
}

func Withdrawal(contract, url, private string, value *big.Int) error {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return errors.New(fmt.Sprintf("send transaction private err:%v", err))
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	from := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("withdrawal", value, "to", from.Hex())
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	instance, err := NewDeposit(common.HexToAddress(contract), client)
	if err != nil {
		return err
	}
	call := &bind.TransactOpts{
		From:  from,
		Nonce: nil,
		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(tx, signer, privateKey)
		},
		Value:    nil,
		GasPrice: nil,
		GasLimit: 3000000,
		Context:  context.Background(),
	}
	transaction, err := instance.Withdrawal(call, value)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", transaction)
	return nil
}
