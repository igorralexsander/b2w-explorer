package service

import (
	"b2w-explorer/internal/app/model"
)

type productPage struct {
	b2wClient Client
}

type ProductPage interface {
	FetchProductPage(request model.ProductPageRequest) (string, error)
}

func NewProductPage(b2wClient Client) *productPage {
	return &productPage{b2wClient}
}

func (s productPage) FetchProductPage(request model.ProductPageRequest) (string, error) {
	url := s.urlBuilder(request)
	response, err := s.b2wClient.FetchPage(url)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (s productPage) urlBuilder(request model.ProductPageRequest) string {
	baseUrl := "https://www."
	switch request.Store {
	case "AMERICANAS":
		baseUrl += "americanas"
	case "SHOPTIME":
		baseUrl += "shoptime"
	case "SUBMARINO":
		baseUrl += "submarino"
	default:
		baseUrl += ""
	}
	baseUrl += ".com.br/produto/"
	baseUrl += request.ProductId
	return baseUrl
}
