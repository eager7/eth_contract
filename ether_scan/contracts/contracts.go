package contracts

import (
	"context"
	"fmt"
	"github.com/eager7/eth_contract/ether_scan/database/tables"
	"github.com/eager7/eth_contract/ether_scan/request"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

type Contract struct {
	hash  string
	count int64
	full  string
	code  string
	abi   string
	bin   string
}

func ParseContracts(ctx context.Context, db *gorm.DB, requester *request.Requester, contracts []*tables.ContractCountInfo, api, dir string) error {
	for _, contract := range contracts {
		c := Contract{hash: contract.Hash, count: contract.Count}
	retry:
		if err := c.RequestContractCode(ctx, db, requester, api, dir); err != nil {
			fmt.Println("get code err:", err)
			time.Sleep(time.Second * 1)
			goto retry
		}
		//if err := db.Where("hash=?", contract.Hash).Delete(&tables.TableContractInfo{}).Error; err != nil && err != gorm.ErrRecordNotFound {
		//	fmt.Println("delete contract err:", err)
		//	goto retry
		//}
	}
	return nil
}

func (c *Contract) RequestContractCode(ctx context.Context, db *gorm.DB, requester *request.Requester, api, dir string) error {
	var contractsInfo []*tables.TableContractInfo
	if err := db.Where("hash=?", c.hash).Find(&contractsInfo).Error; err != nil {
		return err
	}
	if len(contractsInfo) == 0 {
		return nil
	}
	var creator, address string = contractsInfo[0].Creator, contractsInfo[0].Address
	for _, ct := range contractsInfo {
	retry:
		fmt.Println("handle contract:", ct.Address)
		tx, _, err := requester.Client.TransactionByHash(ctx, common.HexToHash(ct.Transaction))
		if err != nil {
			fmt.Println("get tx by hash err:", err)
			time.Sleep(time.Second * 1)
			goto retry
		}
		c.bin = common.Bytes2Hex(tx.Data())
		full, code, abi, _, err := requester.RequestContract(api + common.HexToAddress(ct.Address).Hex())
		if err != nil {
			fmt.Println("get contract err:", err)
			time.Sleep(time.Second * 1)
			goto retry
		}
		if len(full) > len(c.full) {
			c.full = full
		}
		if len(code) > len(c.code) {
			creator, address = ct.Creator, ct.Address
			c.code = code
		}
		if len(abi) > len(c.abi) {
			c.abi = abi
		}

	}
	codeInfo := &tables.TableCodeInfo{
		Creator: creator,
		Address: address,
		Hash:    c.hash,
		Code:    c.code,
		Bin:     c.bin,
		Abi:     c.abi,
	}
	if err := tables.InsertCodeToDB(db, codeInfo); err != nil {
		return err
	}
	//if err := c.StoreContract(fmt.Sprintf("%s/%d_%s/%s-%s/", dir, c.count, c.hash[:10], creator, address)); err != nil {
	//	return err
	//}
	return nil
}

func (c *Contract) StoreContract(dir string) error {
	if err := WriteFile(dir, "full.txt", c.full); err != nil {
		return err
	}
	if err := WriteFile(dir, "abi.bin", c.abi); err != nil {
		return err
	}
	if err := WriteFile(dir, "code.bin", c.bin); err != nil {
		return err
	}
	if err := WriteFile(dir, "code.txt", c.code); err != nil {
		return err
	}
	return nil
}

func WriteFile(dir, name, data string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	file, err := os.Create(dir + name)
	if err != nil {
		return err
	}
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
