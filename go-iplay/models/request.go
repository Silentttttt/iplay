package models

const (
	DefaultPageSize = 10
)

// LoginParams login params
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// IDCardAuthenticationParams IDCardAuthentication params
type IDCardAuthenticationParams struct {
	AuthToken string `json:"auth_token"`
	Name      string `json:"name"`
	IdCard    string `json:"id_card"`
}

type QuizzesListParams struct {
	GameId int64 `json:"game_id"`
}

type PageListParams struct {
	PageNo int `json:"page_no"`
}

type UserQuizzesListParams struct {
	AuthToken string `json:"auth_token"`
	UserId    int64  `json:"user_id"`
	GameId    int64  `json:"game_id"`
}

type DoQuizzesParams struct {
	AuthToken   string `json:"auth_token"`
	UserId      int64  `json:"user_id"`
	GameId      int64  `json:"game_id"`
	QuizzesId   int64  `json:"quizzes_id"`
	ChoiceOptId int64  `json:"choice_opt_id"`
	BetAmount   int64  `json:"bet_amount"`
}

type EndQuizzesParams struct {
	AuthToken   string `json:"auth_token"`
	QuizzesId   int64  `json:"quizzes_id"`
	ChoiceOptId int64  `json:"choice_opt_id"`
}
