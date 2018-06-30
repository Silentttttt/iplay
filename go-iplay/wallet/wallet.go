package wallet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

//SendRawTransaction send raw transaction
func SendRawTransaction(host string, tx *TransactionRequest, data string) (*SendTxResponse, error) {
	params := &rawData{
		data,
	}

	url := host + "/v1/user/rawtransaction"
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
