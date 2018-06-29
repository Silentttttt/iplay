package smartcontract

const (
	contractAddress = "n1vRY4NHBDyQ8hKTXQbbWfPEK5odtRnqRNm"
	adminAddress    = "n1GskdDtrSAaLoR9Beg5sakKfXwyqgPDbft"
	adminPasswd     = "123456"
)

// createGame(payType,type, deadLine, theme, options, amount)
type option struct {
	Odd         float32 `json:"odd"`
	Description string  `json:"description"`
}

type createGameArgs struct {
	payType  uint32    `json:"payType"`
	gameType uint32    `json:"type"`
	deadLine uint64    `json:"deadLine"`
	theme    string    `json:"theme"`
	amount   uint64    `json: "amount"`
	options  []*option `json:"options"`
}
