package controllers

import (
	"github.com/lxndrrud/webviewKsyukulyator/dto"
)

type CalculationController struct {
	products []dto.ProductAmount
	sum      float32
}

func NewCalculationController() *CalculationController {
	calc := &CalculationController{
		products: []dto.ProductAmount{},
		sum:      0,
	}
	calc.GetProducts()
	calc.CalculateSum()
	return calc
}

func (c CalculationController) GetProducts() []dto.ProductAmount {
	return c.products
}

func (c *CalculationController) CalculateSum() float32 {
	c.ClearSum()
	for _, value := range c.products {
		c.sum += value.CalculateAmount()
	}
	return c.sum
}

func (c *CalculationController) AddProductAmount(productAmount dto.ProductAmount) float32 {
	c.products = append(c.products, productAmount)
	return c.CalculateSum()
}

func (c *CalculationController) ClearSum() {
	c.sum = 0
}
