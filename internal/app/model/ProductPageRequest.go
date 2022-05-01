package model

type ProductPageRequest struct {
	Store     string `json:"store"`
	ProductId string `json:"productId"`
}
