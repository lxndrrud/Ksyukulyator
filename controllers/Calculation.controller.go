package controllers

import (
	"github.com/lxndrrud/webviewKsyukulyator/dto"
)

type CalculationController struct {
	products []dto.ProductAmount
	sum      float32
}

func NewCalculationController() *CalculationController {
	return &CalculationController{
		products: []dto.ProductAmount{},
		sum:      0,
	}
}

func (c *CalculationController) GetProducts() []dto.ProductAmount {
	return c.products
}

func (c *CalculationController) CalculateSum() float32 {
	c.sum = 0
	for index := range c.products {
		c.sum += c.products[index].CalculateAmountCost()
	}
	return c.sum
}

func (c *CalculationController) AddProductAmount(productAmount dto.ProductAmount) float32 {
	ptrProduct := &productAmount
	ptrProduct.CalculateAmount()
	isFound := false
	for index, value := range c.products {
		if value.Id == ptrProduct.Id {
			c.products[index].Amount += ptrProduct.Amount
			isFound = true
			break
		}
	}
	if !isFound {
		c.products = append(c.products, productAmount)
	}
	return c.CalculateSum()
}

func (c *CalculationController) DeleteProductAmount(idProduct int64) float32 {
	productsResult := make([]dto.ProductAmount, 0)
	for _, product := range c.products {
		if product.Id != idProduct {
			productsResult = append(productsResult, product)
		}
	}
	c.products = productsResult
	return c.CalculateSum()
}

func (c *CalculationController) ClearProducts() float32 {
	c.sum = 0
	c.products = []dto.ProductAmount{}
	return 0
}
