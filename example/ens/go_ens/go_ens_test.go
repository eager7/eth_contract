package go_ens

import "testing"

const rpc = "https://mainnet.infura.io"
const ws = "wss://mainnet.infura.io/ws"

func TestEns(t *testing.T) {
	ens, err := Initialize(rpc) //"https://infura.io/v3/SECRET"
	if err != nil {
		t.Fatal(err)
	}
	addr, err := ens.EnsAddressByName("wealdtech.eth")
	if err != nil {
		t.Fatal(err)
	}
	addr, err = ens.EnsAddressByName("darkmarket.eth")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addr)

	domain, err := ens.EnsNameByAddress("0x64372db6405879214a0a76a7f1e9c013fd2fd84b")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(domain)
}
