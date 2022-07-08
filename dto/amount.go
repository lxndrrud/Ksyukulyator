package dto

type ProductAmount struct {
	Product
	Amount     float32 `json:"product_amount"`
	AmountCost float32 `json:"product_amount_cost"`
}

func (p *ProductAmount) CalculateAmount() float32 {
	p.AmountCost = p.Product.Cost / 12 * p.Amount / 100
	return p.AmountCost
}
