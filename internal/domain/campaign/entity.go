package campaign

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CampaignStatus string

const (
	Paused  CampaignStatus = "PAUSED"
	Enabled CampaignStatus = "ENABLED"
)

type Campaign struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Status    CampaignStatus `json:"status"`
	CreatedAt time.Time      `json:"createdAt"`
}

func NewCampaign(input CampaignCrate) (*Campaign, error) {
	err := validateCampaign(input)
	if err != nil {
		return nil, err
	}

	return &Campaign{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Status:    input.Status,
		CreatedAt: time.Now(),
	}, nil
}

func validateCampaign(input CampaignCrate) error {
	if input.Name == "" {
		return errors.New("BAD_REQUEST:Campaign name cannot be empty")
	}

	if input.Status == "" {
		return errors.New("BAD_REQUEST:Campaign status cannot be empty")
	}

	if input.Status != Paused && input.Status != Enabled {
		return errors.New("BAD_REQUEST:Campaign status invalid")
	}

	return nil
}
