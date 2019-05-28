package database_test

import (
	"fmt"
	"github.com/eager7/ethereum_contracts/database"
	"testing"
)


func TestSearchContracts(t *testing.T) {
	db, err := database.Initialize("127.0.0.1:3306", "root", "plainchant", "eth_contract", 2, 1)
	if err != nil {
		panic(err)
	}
	contracts, err := database.SearchContracts(db.LogMode(true))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("contracts:", contracts[0])
}

