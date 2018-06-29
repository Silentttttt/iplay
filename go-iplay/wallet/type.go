package wallet

const (
	remoteNebHost = "https://mainnet.nebulas.io"
	localNebHost  = "http://localhost:8685"
)

const (
	gasPrice = "1000000"
	gasLimit = "2000000"
)

//RPCResponse get account state response
type RPCResponse struct {
	// Result *AccountState `json:"result"`
	Result interface{} `json:"result"`
	Err    string      `json:"error"`
}

//GetAccountStateRequest get account state request
type GetAccountStateRequest struct {
	// Hex string of the account addresss.
	Address string `json:"address"`
	// block account state with height. If not specified, use 0 as tail height.
	Height uint64 `json:"height"`
}

//AccountState account state
type AccountState struct {
	// Current balance in unit of 1/(10^18) nas.
	Balance string `json:"balance"`
	// Current transaction count.
	Nonce string `json:"nonce"`
	// Account type
	Type uint32 `json:"type"`
}

//ContractRequest contract request
type ContractRequest struct {
	// call contract function name
	Function string `json:"function"`
	// the params of contract.
	Args string `json:"args"`
}

//TransactionRequest tx request
type TransactionRequest struct {
	// Hex string of the sender account addresss.
	From string `json:"from"`
	// Hex string of the receiver account addresss.
	To string `json:"to"`
	// Amount of value sending with this transaction.
	Value string `json:"value"`
	// Transaction nonce.
	Nonce uint64 `json:"nonce"`
	// gasPrice sending with this transaction.
	GasPrice string `json:"gas_price"`
	// gasLimit sending with this transaction.
	GasLimit string `json:"gas_limit"`
	// contract sending with this transaction
	Contract *ContractRequest `json:"contract"`
	// binary data for transaction
	Binary []byte `json:"binary"`
	// transaction payload type, enum:binary, deploy, call
	Type string `json:"type"`
}

type rawData struct {
	Data string `json:"data"`
}

type addressResponse struct {
	Address string `json:"address"`
}

//SendTxResponse reponse of send tx
type SendTxResponse struct {
	// TxHash tx hash
	TxHash string `json:"txhash"`
	// ContractHash contract hash
	ContractHash string `json:"contract_address"`
}
