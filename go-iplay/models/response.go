package models

type LoginResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data LoginResponseData `json:"data"`
	Ts   int64             `json:"ts"`
}

type LoginResponseData struct {
	AuthToken string
	Username  string
	Avatar    string
}
