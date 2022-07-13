package dto

type ProductAmountTable struct {
	Products []ProductAmount `json:"products"`
	Sum      float32         `json:"sum"`
}
