package tables

import (
	"fmt"
)

type TableContractInfo struct {
	Id          uint64 `json:"id"                gorm:"column:id;primary_key;AUTO_INCREMENT"`      //自增主键
	Address     string `json:"address"           gorm:"column:address;type:char(64);unique_index"` //合约地址
	Creator     string `json:"creator"           gorm:"column:creator;type:char(64)"`              //合约创建者
	Timestamp   uint64 `json:"timestamp"         gorm:"column:timestamp"`                          //合约创建时间
	BlockNumber uint64 `json:"block_number"      gorm:"column:block_number"`                       //创建合约的区块高度
	Transaction string `json:"transaction"       gorm:"column:transaction;type:char(64)"`          //创建合约的交易哈希
	Hash        string `json:"hash"              gorm:"column:hash;type:char(64)"`                 //合约的代码哈希
}

func (t *TableContractInfo) TableName() string {
	return "t_contract_info"
}

func (t *TableContractInfo) GetInstance() interface{} {
	return t
}

func (t *TableContractInfo) SqlCommand() string {
	return fmt.Sprintf("INSERT IGNORE INTO `t_contract_info` (`address`,`creator`,`timestamp`,`block_number`,`transaction`, `hash`) VALUES ('%s','%s','%v','%v','%s', '%s')",
		t.Address, t.Creator, t.Timestamp, t.BlockNumber, t.Transaction, t.Hash)
}

type ContractCountInfo struct {
	Hash  string `json:"hash" gorm:"column:hash;type:char(64)"`
	Count int64  `json:"count" gorm:"column:count"`
}

func (c *ContractCountInfo) TableName() string {
	return "t_contract_count_info"
}