package controllers

import (
	"encoding/json"
	"fmt"
	"iplay/go-iplay/models"
	smartcontract "iplay/go-iplay/smartContract"
	"time"
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
	q.Mapping("create_quizzes", q.CreateQuizzes)
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
// @Success 200
// @Failure 500
// @router /quizzes [post]
func (q *QuizzesController) QuizzesList() {
	var params models.QuizzesListParams
	json.Unmarshal(q.Ctx.Input.RequestBody, &params)
	game, _ := models.GetGameById(params.GameId)
	quizzes, _ := models.GetQuizzesListFromNow(params.GameId)
	gameQuizzesList := &models.GameQuizzesList{Game: game, Quizzes: quizzes}

	q.json(Success, "", gameQuizzesList)
}

// @Title CreateQuizzes
// @Description
// @Success 200
// @Failure 500
// @router /create_quizzes [get]
func (q *QuizzesController) CreateQuizzes() {
	quizzes, _ := models.GetAllQuizzes()
	for k := range *quizzes {
		time, _ := time.Parse("2006-01-02 15:04:05", (*quizzes)[k].End)
		txHash, _ := smartcontract.CreateQuizze(nil, 1, 1, time.Unix()*1000, 1, "1/8决赛", (*quizzes)[k].ChoiceOpt)
		fmt.Println("txhash:", txHash)
	}
	q.json(Success, "", quizzes)
}
