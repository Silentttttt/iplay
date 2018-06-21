package controllers

// FundController fund info
type FundController struct {
	BaseController
}

func (c *FundController) recharge() {
	c.json("", "", nil)
}

func (c *FundController) withdraw() {
	c.json("", "", nil)
}
