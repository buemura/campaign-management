package campaign

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCampaign(input CampaignCrate) *Campaign {
	return &Campaign{
		ID:        uuid.NewString(),
		Name:      input.Name,
		Status:    input.Status,
		CreatedAt: time.Now(),
	}
}
