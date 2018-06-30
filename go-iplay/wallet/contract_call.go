package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"iplay/go-iplay/models"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

//CallContract call contract, if nonce == 0, use get state to get online nonce
func CallContract(
	o orm.Ormer,
	from string,
	to string,
	value string,
	nonce uint64,
	function string,
	params []interface{},
	passwd string) (string, error) {
	//1.get nonce TODO: to use a cache

	result, ok := nonceMgr.Load(from)
	if !ok {
		newGuard := &nonceGuard{}
		result, _ = nonceMgr.LoadOrStore(from, newGuard)
	}
	guard := result.(*nonceGuard)

	guard.lock.Lock()

	if guard.nonce == 0 {
		state, err := GetAccountState(from)
		if err != nil {
			return "", err
		}
		v, err := strconv.Atoi(state.Nonce)
		if err != nil || v < 0 {
			return "", errors.New("failed to get nonce")
		}
		guard.nonce = uint64(v)
	}
	guard.nonce++
	nonce = guard.nonce
	guard.lock.Unlock()

	//2.组装tx
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

	//3.通过本地节点签名
	data, err := SignTxWithPasswd(tx, passwd)
	if err != nil {
		return "", err
	}

	//4.通过send raw tx 到本地节点获得txHash
	var localResp *SendTxResponse
	if localResp, err = SendRawTransaction(localNebHost, tx, data); err != nil {
		return "", err
	}
	fmt.Println(localResp)

	//5. 写数据库记录本次合约调用
	if o != nil {
		record := &models.ContractTransaction{
			From:     from,
			To:       to,
			Value:    value,
			Function: function,
			Args:     string(args),
			Hash:     localResp.TxHash,
			Status:   2, //appending
			Updated:  time.Now(),
		}
		if _, err = o.Insert(record); err != nil {
			return "", nil
		}
	}

	//6. 发送到远程节点
	go func() {
		var remoteResp *SendTxResponse
		for retry := 3; retry > 0; retry-- {
			if remoteResp, err = SendRawTransaction(remoteNebHost, tx, data); err != nil {
				//TODO: use logging
				fmt.Println("faild to send raw tx to remote neb, err", err)
				time.Sleep(time.Millisecond * 100)
			} else {
				break
			}
		}
		if remoteResp.TxHash != localResp.TxHash {
			fmt.Println("Unexpected error: tx different between local and remote")
			fmt.Println(remoteResp)
		}
	}()

	return localResp.TxHash, nil
}
