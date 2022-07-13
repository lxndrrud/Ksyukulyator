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

func (c *CalculationController) GetProducts() dto.ProductAmountTable {
	return dto.ProductAmountTable{
		Products: c.products,
		Sum:      c.sum,
	}
}

func (c *CalculationController) CalculateSum() float32 {
	c.sum = 0
	for index := range c.products {
		c.sum += c.products[index].CalculateAmountCost()
	}
	return c.sum
}

func (c *CalculationController) AddProductAmount(productAmount *dto.ProductAmount) dto.ProductAmountTable {
	productAmount.CalculateAmount()

	c.products = append(c.products, *productAmount)
	return dto.ProductAmountTable{
		Products: c.products,
		Sum:      c.CalculateSum(),
	}
}

func (c *CalculationController) DeleteProductAmount(idProduct int64) dto.ProductAmountTable {
	productsResult := make([]dto.ProductAmount, 0)
	for _, product := range c.products {
		if product.Id != idProduct {
			productsResult = append(productsResult, product)
		}
	}
	c.products = productsResult
	return dto.ProductAmountTable{
		Products: c.products,
		Sum:      c.CalculateSum(),
	}
}

func (c *CalculationController) ClearProducts() dto.ProductAmountTable {
	c.sum = 0
	c.products = []dto.ProductAmount{}
	return dto.ProductAmountTable{
		Products: c.products,
		Sum:      c.sum,
	}
}
