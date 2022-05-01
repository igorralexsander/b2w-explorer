package service

type Client interface {
	FetchPage(url string) ([]byte, error)
}
