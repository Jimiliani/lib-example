package lib_example

import (
	"github.com/Jimiliani/lib-example/internal/clients"
)

type dataFetcher interface {
	clients.Service1Client
}

type Library struct {
	dataFetcher dataFetcher
}

func NewLibrary(dataFetcher dataFetcher) Library {
	return Library{dataFetcher: dataFetcher}
}
