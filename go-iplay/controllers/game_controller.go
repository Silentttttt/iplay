package controllers

import (
	"encoding/json"
	"iplay/go-iplay/models"
)

type GameController struct {
	BaseController
}

type QuizzesController struct {
	BaseController
}

func (game *GameController) URLMapping() {
	game.Mapping("list", game.List)
}

func (q *QuizzesController) URLMapping() {
	q.Mapping("quizzes", q.QuizzesList)
}

// List return game list
// @Title List
// @Description football game list
// @Success 200 {object} models.GameListResponse
// @Failure 500
// @router /list [get]
func (game *GameController) List() {
	games, _ := models.GetGameListFromNow()
	game.json(Success, "", games)
}

// QuizzesList return quizzes list
// @Title QuizzesList
// @Description football game Quizzes
// @Param   data body models.QuizzesListParams true "get quizzes by gameID params"
// @Success 200 {object} models.QuizzesListResponse
// @Failure 500
// @router /quizzes [post]
func (q *QuizzesController) QuizzesList() {
	var params models.QuizzesListParams
	json.Unmarshal(q.Ctx.Input.RequestBody, &params)
	quizzes, _ := models.GetQuizzesListFromNow(params.GameId)
	q.json(Success, "", quizzes)
}
