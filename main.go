package main

import "github.com/eager7/eth_contract/example/ens/go_ens"

func main() {
	_, _ = go_ens.Initialize("https://infura.io/v3/SECRET")
}
