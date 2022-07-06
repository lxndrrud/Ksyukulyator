package dto

type ProductAmount struct {
	Product
	Amount float32
}

func (p ProductAmount) CalculateAmount() float32 {
	return p.Product.Cost / 12 * p.Amount / 100
}
