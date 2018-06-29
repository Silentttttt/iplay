package smartcontract

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

// buyTicket: function(gameId, optionNo, optionVersion, amount) {
type buyTicketArgs struct {
	gameId        uint64 `json:"gameId"`
	optionNo      uint32 `json:"optionNo"`
	optionVersion uint32 `json:"optionVersion"`
	amount        uint64 `json:"amount"`
}

type createGameParams struct {
	functionName string `json:"function"`
	args         string `json:"args"`
}
