package tables

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type TableCodeInfo struct {
	Id      uint64 `json:"id"                gorm:"column:id;primary_key;AUTO_INCREMENT"`      //自增主键
	Creator string `json:"creator"           gorm:"column:creator;type:char(64)"`              //合约创建者
	Address string `json:"address"           gorm:"column:address;type:char(64);unique_index"` //合约地址
	Hash    string `json:"hash"              gorm:"column:hash;type:char(64)"`                 //合约的代码哈希
	Code    string `json:"code"              gorm:"column:code;type:text"`                     //合约的代码
	Bin     string `json:"bin"               gorm:"column:bin;type:text"`                      //合约的代码二进制
	Abi     string `json:"abi"               gorm:"column:abi;type:text"`                      //合约的代码的abi
}

func (t *TableCodeInfo) TableName() string {
	return "t_code_info"
}

func (t *TableCodeInfo) GetInstance() interface{} {
	return t
}

func (t *TableCodeInfo) SqlCommand() string {
	return fmt.Sprintf("INSERT IGNORE INTO `t_code_info` (`creator`,`address`,`hash`,`code`,`bin`, `abi`) VALUES ('%s','%s','%s','%s','%s', '%s')",
		t.Creator, t.Address, t.Hash, t.Code, t.Bin, t.Abi)
}

func InsertCodeToDB(db *gorm.DB, t *TableCodeInfo) error {
	if err := db.Exec(t.SqlCommand()).Error; err != nil {
		return err
	}
	return nil
}