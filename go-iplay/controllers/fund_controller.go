package controllers

// FundController fund info
type FundController struct {
	BaseController
}

func (c *FundController) recharge() {
	c.json(Success, "", nil)
}

func (c *FundController) withdraw() {
	c.json(Success, "", nil)
}
