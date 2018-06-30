package models

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}
type LoginResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data LoginResponseData `json:"data"`
	Ts   int64             `json:"ts"`
}

type LoginResponseData struct {
	AuthToken   string `json:"auth_token"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	HashAddress string `json:"hash_address"`
	Balance     int64  `json:"balance"`
}

type GameListResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Game `json:"data"`
	Ts   int64  `json:"ts"`
}

type QuizzesListResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data GameQuizzesList `json:"data"`
	Ts   int64           `json:"ts"`
}

type GameQuizzesList struct {
	Game    *Game      `json:"game"`
	Quizzes *[]Quizzes `json:"quizzes"`
}

type UserQuizzesListResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []UserQuizzes `json:"data"`
	Ts   int64         `json:"ts"`
}
