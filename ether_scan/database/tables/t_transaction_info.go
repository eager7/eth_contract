package tables

import (
	"fmt"
)

type TableTransactionInfo struct {
	Id               uint64 `json:"id"                gorm:"column:id;primary_key;AUTO_INCREMENT"` //自增主键
	BlockHash        string `json:"block_hash"        gorm:"column:block_hash;type:char(64)"`      //当前交易所在区块的哈希
	BlockNumber      uint64 `json:"block_number"      gorm:"column:block_number"`                  //当前交易所在区块的高度
	Timestamp        uint64 `json:"timestamp"         gorm:"column:timestamp"`                     //交易发生的时间点
	From             string `json:"from"              gorm:"column:from;type:char(64)"`            //当前交易发起者
	To               string `json:"to"                gorm:"column:to;type:char(64)"`              //当前交易接收者
	Gas              uint64 `json:"gas"               gorm:"column:gas"`                           //当前交易给定gas量
	GasUsed          uint64 `json:"gas_used"          gorm:"column:gas_used"`                      //当前交易使用gas量
	GasPrice         string `json:"gas_price"         gorm:"column:gas_price"`                     //当前交易gas单价
	Hash             string `json:"hash"              gorm:"column:hash;type:char(64)"`            //当前交易哈希值
	InputFlag        uint8  `json:"input_flag"        gorm:"column:input_flag"`                    //标识input数据位置，1表示从数据库取，2表示从链上取
	Status           uint64 `json:"status"            gorm:"column:status"`                        //标识交易执行情况，1表示执行成功，其他的从黄皮书查询错误码
	Input            string `json:"input"             gorm:"column:input"`                         //交易输入数据，如果to是合约，这里存放合约调用方法和参数，最大长度15360字节
	Nonce            uint64 `json:"nonce"             gorm:"column:nonce"`                         //交易记数，是一个递增的值，防止交易重放
	TransactionIndex uint16 `json:"transaction_index" gorm:"column:transaction_index"`             //交易在区块中的偏移量
	Value            string `json:"value"             gorm:"column:value"`                         //交易金额
	V                string `json:"v"                 gorm:"column:v"`                             //交易签名
	S                string `json:"s"                 gorm:"column:s"`                             //交易签名
	R                string `json:"r"                 gorm:"column:r"`                             //交易签名
}

func (t *TableTransactionInfo) TableName() string {
	return "t_transaction_info"
}

func (t *TableTransactionInfo) GetInstance() interface{} {
	return t
}

func (t *TableTransactionInfo) SqlCommand() string {
	return fmt.Sprintf("INSERT IGNORE INTO `t_transaction_info` "+
		"(`block_hash`,`block_number`,`timestamp`,`from`,`to`,`gas`,`gas_used`,`gas_price`,`hash`,`input_flag`,`input`,`nonce`,`transaction_index`,`value`,`status`,`v`,`s`,`r`) "+
		"VALUES ('%s','%v','%v','%s','%s','%v','%v','%s','%s','%v','%s','%v','%v','%s','%v','%s','%s','%s') ",
		t.BlockHash, t.BlockNumber, t.Timestamp, t.From, t.To, t.Gas, t.GasUsed, t.GasPrice, t.Hash, t.InputFlag, t.Input, t.Nonce, t.TransactionIndex, t.Value, t.Status, t.V, t.S, t.R)
}
