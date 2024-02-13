package lib_example

import (
	"github.com/Jimiliani/lib-example/internal/models"
)

type dataFetcher interface {
	GetProcessedData() (*models.ProcessedData, error)
}

type Library struct {
	dataFetcher dataFetcher
}

func NewLibrary(dataFetcher dataFetcher) Library {
	return Library{dataFetcher: dataFetcher}
}
