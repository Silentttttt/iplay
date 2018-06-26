package controllers

import (
	"encoding/json"
	"iplay/go-iplay/models"
)

type UserQuizzesController struct {
	BaseController
}

func (uq *UserQuizzesController) URLMapping() {
	uq.Mapping("quizzes_list", uq.UserQuizzesList)
}

type UserQuizzesListParams struct {
	UserId int64
}

// QuizzesList return user quizzes list
// @Title UserQuizzesList
// @Description football game user quizzes
// @Param   data body controllers.UserQuizzesListParams true "get user quizzes request params"
// @Success 200 {object} models.UserQuizzesListResponse
// @Failure 500
// @router /quizzes [post]
func (uq *UserQuizzesController) UserQuizzesList() {
	var params UserQuizzesListParams
	json.Unmarshal(uq.Ctx.Input.RequestBody, &params)
	userQuizzes, _ := models.GetUserQuizzesList(params.UserId)
	uq.json(Success, "", userQuizzes)
}
