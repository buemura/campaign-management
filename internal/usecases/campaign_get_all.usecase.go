package usecases

import "github.com/bemura/campaign-management/internal/domain/campaign"

type CampaignGetAllUsecase struct {
	campaignRepository campaign.CampaignRepository
}

func NewCampaignGetAllUsecase(campaignRepository campaign.CampaignRepository) *CampaignGetAllUsecase {
	return &CampaignGetAllUsecase{
		campaignRepository: campaignRepository,
	}
}

func (u *CampaignGetAllUsecase) Execute() ([]*campaign.Campaign, error) {
	c, err := u.campaignRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return c, nil
}
