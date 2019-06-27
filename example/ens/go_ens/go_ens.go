package go_ens

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens"
)

type Ens struct {
	client *ethclient.Client
}

func Initialize(url string) (*Ens, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, errors.New("dial eth client err:" + err.Error())
	}
	return &Ens{client: client}, nil
}

func (e *Ens) EnsNameByAddress(addr string) (string, error) {
	address := common.HexToAddress(addr)
	reverse, err := ens.ReverseResolve(e.client, &address)
	if err != nil {
		return "", errors.New("resolve addr err:" + err.Error())
	}
	if reverse == "" {
		fmt.Printf("%s has no reverse lookup\n", address.Hex())
	} else {
		fmt.Printf("Name of %s is %s\n", address.Hex(), reverse)
	}
	return reverse, nil
}

func (e *Ens) EnsAddressByName(domain string) (string, error) {
	address, err := ens.Resolve(e.client, domain)
	if err != nil {
		return "", errors.New("resolve domain err:" + err.Error())
	}
	fmt.Printf("Address of %s is %s\n", domain, address.Hex())
	return address.Hex(), nil
}
