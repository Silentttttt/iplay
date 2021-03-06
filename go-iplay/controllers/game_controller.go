package controllers

import (
	"encoding/json"
	"fmt"
	"iplay/go-iplay/models"
	smartcontract "iplay/go-iplay/smartContract"
	"time"
	t "time"
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
// @Param data body models.PageListParams true "get game list by pager"
// @Success 200 {object} models.GameListResponse
// @Failure 500
// @router /list [post]
func (game *GameController) List() {
	var params models.PageListParams
	json.Unmarshal(game.Ctx.Input.RequestBody, &params)
	page, _ := models.GetGameListFromNow(params.PageNo)
	game.json(Success, "", page)
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
		end, _ := time.Parse("2006-01-02 15:04:05", (*quizzes)[k].End)
		txHash, _ := smartcontract.CreateQuizze(nil, 1, 1, end.Unix()*1000, 1, "1/8决赛", (*quizzes)[k].ChoiceOpt)
		fmt.Println("txhash:", txHash)
		t.Sleep(2 * time.Second)
	}
	q.json(Success, "", quizzes)
}
