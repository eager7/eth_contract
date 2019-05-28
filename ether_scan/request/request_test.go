package request

import (
	"fmt"
	"testing"
)

func TestRequestContract(t *testing.T) {
	c, err := (&Requester{}).RequestContract("https://etherscan.io/address/0x4204B84f008f469389698c622F18Ed2c9bc74D7F")
	//_, _, _, err := (&Requester{}).RequestContract("https://etherscan.io/address/0x5e7Bb4b74480738871564d1e75Dd5B183F9892A8")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(c)
}

func TestRequestIpList(t *testing.T) {
	_, err := GetIpList()
	if err != nil {
		t.Fatal(err)
	}
}