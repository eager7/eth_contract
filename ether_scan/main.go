package main

import (
	"context"
	"fmt"
	"github.com/eager7/eth_contract/ether_scan/config"
	"github.com/eager7/eth_contract/ether_scan/contracts"
	"github.com/eager7/eth_contract/ether_scan/database"
	"github.com/eager7/eth_contract/ether_scan/request"
	"sync"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}
	db, err := database.Initialize(cfg.DbOpt.Address, cfg.DbOpt.User, cfg.DbOpt.Password, cfg.DbOpt.DbName, cfg.DbOpt.MaxOpenConn, cfg.DbOpt.MaxIdleConn)
	if err != nil {
		panic(err)
	}
	requester, err := request.Initialize(cfg.EthOpt.Address)
	if err != nil {
		panic(err)
	}
	cons, err := database.SearchContracts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("search contracts len:", len(cons))
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(cfg.Worker)
	index := len(cons) / cfg.Worker
	for i := 0; i < cfg.Worker; i++ {
		go func(offset int) {
			start := index * offset
			end := index * (offset + 1)
			if offset == cfg.Worker-1 {
				end = len(cons)
			}
			fmt.Println("start:", start, "to:", end)
			if err := contracts.ParseContracts(ctx, db, requester, cons[start:end], cfg.EthOpt.ApiAddress, cfg.Path); err != nil {
				panic(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	cancel()
}
