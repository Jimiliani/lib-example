package clients

import "net/http"

type Service1Client interface {
	GetSomeData() (*SomeData, error)
}

type service1Client struct {
	httpClient *http.Client
}

type SomeData struct {
	ID int64
}

func NewService1Client() Service1Client {
	return &service1Client{
		httpClient: &http.Client{},
	}
}

func (m *service1Client) GetSomeData() (*SomeData, error) {
	// Use m.httpClient to make HTTP requests
	return &SomeData{}, nil
}
