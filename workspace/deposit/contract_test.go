package deposit_test

import (
	"github.com/eager7/eth_contract/workspace/deposit"
	"math/big"
	"testing"
)

func TestWithdrawal(t *testing.T) {
	contract := "0x1732899509c0239830b6fcec92d85df84e66b2a9"
	private := "d56daa057596332bf93ebc2159e86ad284a69d73223b5e1982c61faf60295c1c"
	if err := deposit.Withdrawal(contract, "http://127.0.0.1:7545", private, new(big.Int).SetUint64(200)); err != nil {
		t.Fatal(err)
	}
}
