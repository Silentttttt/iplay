package wallet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func sendRPC(method string, url string, buf []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(buf))
	if err != nil {
		//TODO:
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		//TODO:
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //TODO: 设置超时时间
	if err != nil {
		//TODO:
		return nil, err
	}

	return body, nil
}

//GetAccountState get account state
func GetAccountState(address string) (*AccountState, error) {
	// curl -i -H 'Content-Type: application/json' -X POST http://localhost:8685/v1/user/accountstate -d '{"address":"n1Z6SbjLuAEXfhX1UJvXT6BB5osWYxVg3F3"}'
	url := remoteNebHost + "/v1/user/accountstate"
	params, err := json.Marshal(&GetAccountStateRequest{Address: address})
	if err != nil {
		//TODO:
		return nil, err
	}
	data, err := sendRPC("POST", url, params)
	if err != nil {
		//TODO:
		return nil, err
	}

	rpcResponse := RPCResponse{Result: &AccountState{}}
	if err = json.Unmarshal(data, &rpcResponse); err != nil {
		//TODO:
		fmt.Println(err)
		return nil, err
	}
	if rpcResponse.Err == "" {

		fmt.Println("=============")
	}

	if rpcResponse.Err != "" {
		return nil, errors.New(rpcResponse.Err)
	}

	return rpcResponse.Result.(*AccountState), nil
}

//SignTxWithPasswd sign tx with passwd
func SignTxWithPasswd(tx *TransactionRequest, passwd string) (string, error) {
	params := struct {
		Tx         *TransactionRequest `json:"transaction"`
		Passphrase string              `json:"passphrase"`
	}{
		tx,
		passwd,
	}

	url := localNebHost + "/v1/admin/sign"
	paramsBuf, err := json.Marshal(params)
	if err != nil {
		//TODO:
		return "", err
	}

	data, err := sendRPC("POST", url, paramsBuf)
	if err != nil {
		return "", err
	}

	rpcResponse := RPCResponse{Result: &rawData{}}
	if err = json.Unmarshal(data, &rpcResponse); err != nil {
		return "", err
	}

	if rpcResponse.Err != "" {
		return "", errors.New(rpcResponse.Err)
	}

	return rpcResponse.Result.(*rawData).Data, nil

	//curl -i -H 'Content-Type: application/json' -X POST http://localhost:8685/v1/admin/sign -d
	//'{"transaction":{"from":"n1QZMXSZtW7BUerroSms4axNfyBGyFGkrh5",
	//"to":"n1QZMXSZtW7BUerroSms4axNfyBGyFGkrh5", "value":"1000000000000000000","nonce":1,"gasPrice":"1000000","gasLimit":"2000000"}, "passphrase":"passphrase"}'

}

// curl -i -H 'Content-Type: application/json' -X POST http://localhost:8685/v1/admin/account/new -d '{"passphrase":"passphrase"}'

//CreateAccount create account
func CreateAccount(passwd string) (string, error) {
	params := struct {
		Passwd string `json:"passphrase"`
	}{
		passwd,
	}

	url := localNebHost + "/v1/admin/account/new"
	paramsBuf, err := json.Marshal(params)
	if err != nil {
		//TODO:
		return "", err
	}

	data, err := sendRPC("POST", url, paramsBuf)
	if err != nil {
		return "", err
	}

	rpcResponse := RPCResponse{Result: &addressResponse{}}
	if err = json.Unmarshal(data, &rpcResponse); err != nil {
		return "", err
	}

	if rpcResponse.Err != "" {
		return "", errors.New(rpcResponse.Err)
	}

	return rpcResponse.Result.(*addressResponse).Address, nil
}

//SendTransactionWithPasswd send tx with passwd
func SendTransactionWithPasswd(tx *TransactionRequest, passwd string) (*SendTxResponse, error) {
	data, err := SignTxWithPasswd(tx, passwd)
	if err != nil {
		return nil, err
	}
	params := &rawData{
		data,
	}

	url := remoteNebHost + "/v1/user/rawtransaction"
	paramsBuf, err := json.Marshal(params)
	if err != nil {
		//TODO:
		return nil, err
	}

	dataBuf, err := sendRPC("POST", url, paramsBuf)
	if err != nil {
		return nil, err
	}

	rpcResponse := RPCResponse{Result: &SendTxResponse{}}
	if err = json.Unmarshal(dataBuf, &rpcResponse); err != nil {
		return nil, err
	}

	if rpcResponse.Err != "" {
		return nil, errors.New(rpcResponse.Err)
	}

	return rpcResponse.Result.(*SendTxResponse), nil
}

//CallContract call contract, if nonce == 0, use get state to get online nonce
func CallContract(
	from string,
	to string,
	value string,
	nonce uint64,
	function string,
	params []interface{},
	passwd string) (string, error) {

	if nonce == 0 {
		state, err := GetAccountState(from)
		if err != nil {
			return "", err
		}
		v, err := strconv.Atoi(state.Nonce)
		if err != nil || v < 0 {
			return "", errors.New("failed to get nonce")
		}
		nonce = uint64(v)
		nonce++
	}

	args, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	contract := &ContractRequest{
		Function: function,
		Args:     string(args),
	}

	tx := &TransactionRequest{
		From:     from,
		To:       to,
		Value:    value,
		Nonce:    nonce,
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Contract: contract,
	}
	sendTxResp, err := SendTransactionWithPasswd(tx, passwd)
	if err != nil {
		return "", err
	}
	fmt.Println(sendTxResp)
	return sendTxResp.TxHash, nil
}
