package wallet

import (
	"bytes"
	"encoding/json"
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//TODO:
		return nil, err
	}

	return body, nil
}

func getAccountState(address string) (*AccountState, error) {
	// curl -i -H 'Content-Type: application/json' -X POST http://localhost:8685/v1/user/accountstate -d '{"address":"n1Z6SbjLuAEXfhX1UJvXT6BB5osWYxVg3F3"}'
	url := hostname + "/v1/user/accountstate"
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
	fmt.Println(string(data))

	rpcResponce := RPCResponse{&AccountState{}}
	if err = json.Unmarshal(data, &rpcResponce); err != nil {
		//TODO:
		fmt.Println(err)
		return nil, err
	}
	return rpcResponce.Result.(*AccountState), nil
}

// func signTxWithPasswd() (string, error) {
// }

// func sendTransaction() (string, error) {

// }
