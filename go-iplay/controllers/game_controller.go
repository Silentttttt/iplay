package controllers

import "iplay/go-iplay/models"

type GameController struct {
	BaseController
}

func (game *GameController) URLMapping() {
	game.Mapping("list", game.List)
}

// List return game list
// @Title List
// @Description football game list
// @Success 200 {object} models.GameListResponse
// @Failure 500
// @router /list [post]
func (game *GameController) List() {
	games, _ := models.GetGameListFromNow()
	game.json(Success, "", games)
}
