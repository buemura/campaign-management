package campaign

type CampaignCrate struct {
	Name   string         `json:"name"`
	Status CampaignStatus `json:"status"`
}
