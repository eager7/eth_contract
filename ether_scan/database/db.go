package database

import (
	"fmt"
	"github.com/eager7/eth_contract/ether_scan/database/tables"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Initialize(addr, user, password, dbName string, maxOpenConns, maxIdleConns int) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, password, addr, dbName)
	gdb, err := gorm.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	gdb.DB().SetMaxOpenConns(maxOpenConns)
	gdb.DB().SetMaxIdleConns(maxIdleConns)

	return gdb, err
}

func SearchContracts(db *gorm.DB) ([]*tables.ContractCountInfo, error) {
	var contracts []*tables.ContractCountInfo
	err := db.Table("t_contract_info").Select("hash, count(`hash`) as count").Group("hash").Order("count desc").Limit(10).Scan(&contracts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return contracts, nil
}
