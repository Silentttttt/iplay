package controllers

import (
	"encoding/json"
	"iplay/go-iplay/models"
	"iplay/go-iplay/smartContract"
	"iplay/go-iplay/utils"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
)

type UserQuizzesController struct {
	BaseController
}

func (uq *UserQuizzesController) URLMapping() {
	uq.Mapping("quizzes_list", uq.UserQuizzesList)
	uq.Mapping("do_quizzes", uq.DoQuizzes)
	uq.Mapping("ui_h*9yh/end_quizzes", uq.EndQuizzes)
}

// UserQuizzesList return user quizzes list
// @Title UserQuizzesList
// @Description football game user quizzes
// @Param   data body models.UserQuizzesListParams true "get user quizzes request params"
// @Success 200
// @Failure 500
// @router /quizzes_list [post]
func (uq *UserQuizzesController) UserQuizzesList() {
	var params models.UserQuizzesListParams
	json.Unmarshal(uq.Ctx.Input.RequestBody, &params)
	if CheckAuthToken(params.AuthToken) {

		// userQuizzes, _ := models.GetUserQuizzesList(params.UserId)

		// uq.json(Success, "", userQuizzes)
		var userQuizzesDataList []*models.UserQuizzesData
		games, _ := models.GetUserQuizzesGroupByGame(params.UserId, params.GameId)
		for k := range *games {
			game, _ := models.GetGameById((*games)[k].GameId)
			userQuizzes, _ := models.GetUserQuizzesListByGame(params.UserId, (*games)[k].GameId)
			userQuizzesData := &models.UserQuizzesData{Game: game, UserQuizzes: userQuizzes}
			userQuizzesDataList = append(userQuizzesDataList, userQuizzesData)
		}
		uq.json(Success, "", userQuizzesDataList)

	} else {
		uq.json(NeedLogin, NeedLoginErr, nil)
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
		m := models.UserQuizzes{}
		o := orm.NewOrm()
		user, err := models.GetUserById(params.UserId)
		if err != nil || user == nil {
			uq.json(Fail, "", nil)
			return
		}
		choiceOpt, err := models.GetChoiceOptById(params.ChoiceOptId)
		if err != nil || choiceOpt == nil {
			uq.json(Fail, "", nil)
			return
		}
		game, err := models.GetGameById(params.GameId)
		if err != nil || game == nil {
			uq.json(Fail, "", nil)
			return
		}
		quizzes, err := models.GetQuizzesById(params.QuizzesId)
		if err != nil || quizzes == nil {
			uq.json(Fail, "", nil)
			return
		}
		m.User = user
		m.ChoiceOpt = choiceOpt
		m.Quizzes = quizzes
		m.Money = params.BetAmount
		m.Result = 0
		m.Reward = int64(float32(params.BetAmount) * choiceOpt.Odds)
		m.Game = game
		m.Created = time.Now().Format("2006-01-02 15:04:05")
		o.Begin()
		if _, err := o.Insert(&m); err != nil {
			uq.json(Fail, "", nil)
			return
		}
		user.Balance -= params.BetAmount
		if _, err := o.Update(user); err != nil {
			o.Rollback()
			logs.Error("[DoQuizzes] Failed to update user balance, ", err)
			uq.json(Fail, "", nil)
			return
		}
		_, err = smartcontract.BuyTicket(o, user.HashAddress, user.Passphrase, uint64(game.Id), uint8(choiceOpt.Id), 1, uint64(m.Money))
		if err != nil {
			logs.Error("[DoQuizzes] Failed to do quizzes on chain, ", err)
			o.Rollback()
			uq.json(Fail, "", nil)
			return
		}
		o.Commit()
		uq.json(Success, "", user)
	} else {
		uq.json(NeedLogin, NeedLoginErr, nil)
	}

}

// EndQuizzes do quizzes
// @Title EndQuizzes
// @Description api for ending quizzes
// @Param   data body models.EndQuizzesParams true "params for ending quizzes"
// @Success 200
// @Failure 500
// @router /ui_h*9yh/end_quizzes [post]
func (uq *UserQuizzesController) EndQuizzes() {
	var params models.EndQuizzesParams
	json.Unmarshal(uq.Ctx.Input.RequestBody, &params)
	if CheckAuthToken(params.AuthToken) {
		if utils.Get(params.AuthToken) != "leon" {
			uq.json(Fail, "permission denied ", nil)
			return
		}
		userQuizzes, err := models.GetUserQuizzesByQuizzesAndChoiceOpt(params.QuizzesId, params.ChoiceOptId)
		if err != nil {
			logs.Error("")
			uq.json(Fail, "", nil)
			return
		}
		o := orm.NewOrm()
		o.Begin()
		for k := range *userQuizzes {
			singleUserQuizzes := (*userQuizzes)[k]
			user := singleUserQuizzes.User
			user.Balance += singleUserQuizzes.Reward
			if _, err := o.Update(user); err != nil {
				o.Rollback()
				logs.Error("[EndQuizzes] Failed to update user balance, ", err)
				uq.json(Fail, "", nil)
				return
			}
			singleUserQuizzes.Status = 1
			singleUserQuizzes.Result = 1
			if _, err := o.Update(&singleUserQuizzes); err != nil {
				o.Rollback()
				logs.Error("[EndQuizzes] Failed to update user quizzes status, ", err)
				uq.json(Fail, "", nil)
				return
			}
		}
		o.Commit()
		uq.json(Success, "", nil)
		return
	}

	uq.json(NeedLogin, NeedLoginErr, nil)
}
