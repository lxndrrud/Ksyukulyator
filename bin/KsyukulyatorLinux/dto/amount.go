package dto

type ProductAmount struct {
	Product
	Amount     float32 `json:"product_amount"`
	AmountCost float32 `json:"product_amount_cost"`
}

func (p *ProductAmount) CalculateAmountCost() float32 {
	p.AmountCost = p.Product.Cost / 12 * p.Amount / 100
	return p.AmountCost
}

func (p *ProductAmount) CalculateAmount() float32 {
	if p.AmountCost != 0 && p.Product.Cost != 0 {
		p.Amount = p.AmountCost * 12 * 100 / p.Product.Cost
		return p.Amount
	}
	return p.Amount
}
