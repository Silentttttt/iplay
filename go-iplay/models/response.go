package models

type LoginResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data LoginResponseData `json:"data"`
	Ts   int64             `json:"ts"`
}

type LoginResponseData struct {
	AuthToken string `json:"auth_token"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
}

type GameListResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []Game `json:"data"`
	Ts   int64  `json:"ts"`
}

type QuizzesListResponse struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Quizzes `json:"data"`
	Ts   int64   `json:"ts"`
}

type UserQuizzesListResponse struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []UserQuizzes `json:"data"`
	Ts   int64         `json:"ts"`
}
