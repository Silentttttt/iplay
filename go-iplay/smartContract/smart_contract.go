package smartcontract

import (
	"encoding/json"
	"fmt"

	"iplay/go-iplay/models"
	"iplay/go-iplay/wallet"
)

func createQuizze(
	payType uint32,
	gameType uint32,
	deadLine int64,
	amount uint64,
	theme string,
	opts []models.ChoiceOpt) (string, error) {
	smartContractOpt := make([]*option, 0)

	opt1 := option{1, "fs"}
	a, _ := json.Marshal(opt1)
	fmt.Println(string(a))
	for _, opt := range opts {
		smartContractOpt = append(smartContractOpt, &option{opt.Odds, opt.Name})
	}
	params := make([]interface{}, 0)
	params = append(params, payType)
	params = append(params, gameType)
	params = append(params, deadLine)
	params = append(params, theme)
	params = append(params, smartContractOpt)
	params = append(params, amount)

	txHash, err := wallet.CallContract(adminAddress, contractAddress, "0", 0, "createAndStartGame", params, adminPasswd)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

//TODO: 将参数打包的逻辑封装到callContract中
// Transfer claim token
func Transfer(to string, amount uint64) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, to)
	params = append(params, amount)

	txHash, err := wallet.CallContract(adminAddress, contractAddress, "0", 0, "transfer", params, adminPasswd)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

// buyTicket: function(gameId, optionNo, optionVersion, amount)
//BuyTicket buy ticket
func BuyTicket(buyer string, passwd string, gameID uint64, optionNo uint8, optionVersion uint8, amount uint64) (string, error) {
	params := make([]interface{}, 0)
	params = append(params, gameID)
	params = append(params, optionNo)
	params = append(params, optionVersion)
	params = append(params, amount)

	txHash, err := wallet.CallContract(buyer, contractAddress, "0", 0, "buyTicket", params, passwd)
	if err != nil {
		return "", err
	}
	return txHash, nil
}
