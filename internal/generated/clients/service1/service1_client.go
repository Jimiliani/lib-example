package service1

import "net/http"

type Client interface {
	GetSomeData() (*SomeData, error)
}

type client struct {
	httpClient *http.Client
}

type SomeData struct {
	ID int64
}

func NewService1Client() Client {
	return &client{
		httpClient: &http.Client{},
	}
}

func (m *client) GetSomeData() (*SomeData, error) {
	// Use m.httpClient to make HTTP requests
	return &SomeData{}, nil
}
