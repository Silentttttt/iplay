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
	uq.Mapping("do_quizzes", uq.DoQuizzes)
}

// UserQuizzesList return user quizzes list
// @Title UserQuizzesList
// @Description football game user quizzes
// @Param   data body models.UserQuizzesListParams true "get user quizzes request params"
// @Success 200 {object} models.UserQuizzesListResponse
// @Failure 500
// @router /quizzes_list [post]
func (uq *UserQuizzesController) UserQuizzesList() {
	var params models.UserQuizzesListParams
	json.Unmarshal(uq.Ctx.Input.RequestBody, &params)
	if CheckAuthToken(params.AuthToken) {
		userQuizzes, _ := models.GetUserQuizzesList(params.UserId)
		uq.json(Success, "", userQuizzes)
	} else {
		uq.json(Fail, NeedLoginErr, nil)
	}

}

// DoQuizzes do quizzes
// @Title DoQuizzes
// @Description api for user doing quizzes
// @Param   data body models.DoQuizzesParams true "params for user doing quizzes"
// @Success 200
// @Failure 500
// @router /do_quizzes [post]
func (uq *UserQuizzesController) DoQuizzes() {
	var params models.DoQuizzesParams
	json.Unmarshal(uq.Ctx.Input.RequestBody, &params)
	if CheckAuthToken(params.AuthToken) {
		if err := models.SaveUserQuizzes(&params); err != nil {
			uq.json(Fail, "", nil)
		} else {
			uq.json(Success, "", nil)
		}
	} else {
		uq.json(Fail, NeedLoginErr, nil)
	}

}
