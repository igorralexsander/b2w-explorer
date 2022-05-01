package clients

import (
	"b2w-explorer/internal/app/service"
	"github.com/igorralexsander/httpcircuited"
)

type b2wClient struct {
	httpcircuited.Downstream
}

func NewB2WClient(downstream httpcircuited.Downstream) service.Client {
	return &b2wClient{downstream}
}

func (c b2wClient) FetchPage(url string) ([]byte, error) {
	response, err := c.Downstream.Get(url)
	if response == nil {
		return nil, err
	}
	return response, nil
}
