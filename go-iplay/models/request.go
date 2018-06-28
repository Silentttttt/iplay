package models

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

type UserQuizzesListParams struct {
	AuthToken string `json:"auth_token"`
	UserId    int64  `json:"user_id"`
}

type DoQuizzesParams struct {
	AuthToken   string  `json:"auth_token"`
	UserId      int64   `json:"user_id"`
	QuizzesId   int64   `json:"quizzes_id"`
	ChoiceOptId int64   `json:"choice_opt_id"`
	BetAmount   float64 `json:"bet_amount"`
}