package wallet

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAccountState(t *testing.T) {
	state, err := GetAccountState("n1GskdDtrSAaLoR9Beg5sakKfXwyqgPDbft")
	assert.Equal(t, nil, err)
	fmt.Println(state)

	state, err = GetAccountState("n1dZZnqKGEkb1LHYunTio1j1q")
	assert.Equal(t, errors.New("address: invalid address format"), err)
}

func TestCreateAccount(t *testing.T) {
	address, err := CreateAccount("123456")
	assert.Nil(t, err)
	fmt.Println(address)
}

func TestSignTxWithPasswd(t *testing.T) {
	passwd := "123456"
	address, err := CreateAccount(passwd)
	assert.Nil(t, err)
	fmt.Println(address)

	tx := &TransactionRequest{
		From:     address,
		To:       address,
		Value:    "0",
		Nonce:    0,
		GasPrice: "1000",
		GasLimit: "10000",
		Contract: nil,
		Binary:   nil,
		Type:     "",
	}
	data, err := SignTxWithPasswd(tx, passwd)
	assert.Nil(t, err)
	fmt.Println(data)
}

// func TestSendTx(t *testing.T) {
// 	passwd := "123456"
// 	address := "n1GskdDtrSAaLoR9Beg5sakKfXwyqgPDbft"
// 	tx := &TransactionRequest{
// 		From:     address,
// 		To:       address,
// 		Value:    "1000000",
// 		Nonce:    1,
// 		GasPrice: "1000000",
// 		GasLimit: "2000000",
// 		Contract: nil,
// 		Binary:   nil,
// 		Type:     "",
// 	}
// 	sendTxResp, err := SendTransactionWithPasswd(tx, passwd)
// 	assert.Nil(t, err)
// 	fmt.Println(sendTxResp)
// }
