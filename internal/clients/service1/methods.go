package service1

import (
	"fmt"

	"github.com/Jimiliani/lib-example/internal/models"
)

func (c *Client) GetProcessedData() (*models.ProcessedData, error) {
	data, err := c.client.GetSomeData()
	if err != nil {
		return nil, fmt.Errorf("GetProcessedData error: %w", err)
	}
	if data == nil {
		return nil, fmt.Errorf("GetProcessedData got empty data")
	}
	return &models.ProcessedData{ID: data.ID}, nil
}
