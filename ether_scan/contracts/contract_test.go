package contracts

import (
	"context"
	"fmt"
	"github.com/eager7/ethereum_contracts/config"
	"github.com/eager7/ethereum_contracts/database"
	"github.com/eager7/ethereum_contracts/database/tables"
	"github.com/eager7/ethereum_contracts/request"
	"testing"
)

func TestWriteFile(t *testing.T) {
	if err := WriteFile("/tmp/test_dir/", "code.bin", "test"); err != nil {
		t.Fatal(err)
	}
}

func TestParseContracts(t *testing.T) {
	cfg, err := config.Initialize()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.Initialize(cfg.DbOpt.Address, cfg.DbOpt.User, cfg.DbOpt.Password, cfg.DbOpt.DbName, cfg.DbOpt.MaxOpenConn, cfg.DbOpt.MaxIdleConn)
	if err != nil {
		t.Fatal(err)
	}
	requester, err := request.Initialize(cfg.EthOpt.Address)
	if err != nil {
		t.Fatal(err)
	}

	c := tables.ContractCountInfo{Hash: "67d52d39181d7e8e8b3342808ba64cab067564ed4397bd36d4f80ce7a075f774", Count: 92}
	if err := ParseContracts(context.Background(), db.LogMode(true),requester,  []*tables.ContractCountInfo{&c,},cfg.EthOpt.ApiAddress, "/tmp/test_my"); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", c)
}
