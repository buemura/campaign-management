package campaign

type CampaignCrate struct {
	Name   string         `json:"name"`
	Status CampaignStatus `json:"status"`
}

type CampaignUpdate struct {
	ID     string          `json:"id"`
	Name   *string         `json:"name"`
	Status *CampaignStatus `json:"status"`
}
