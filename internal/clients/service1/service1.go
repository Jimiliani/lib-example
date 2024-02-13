package service1

import (
	"github.com/Jimiliani/lib-example/internal/generated/clients/service1"
)

type Client struct {
	client service1.Client
}

func New() Client {
	return Client{client: service1.NewService1Client()}
}
