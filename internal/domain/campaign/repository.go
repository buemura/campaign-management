package campaign

type CampaignRepository interface {
	FindAll() ([]*Campaign, error)
	FindById(id string) (*Campaign, error)
	Save(input *Campaign) (*Campaign, error)
	Delete(id string) error
}
