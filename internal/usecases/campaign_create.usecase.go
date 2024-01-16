package usecases

import "github.com/bemura/campaign-management/internal/domain/campaign"

type CampaignCreateUsecase struct {
	campaignRepository campaign.CampaignRepository
}

func NewCampaignCreateUsecase(campaignRepository campaign.CampaignRepository) *CampaignCreateUsecase {
	return &CampaignCreateUsecase{
		campaignRepository: campaignRepository,
	}
}

func (u *CampaignCreateUsecase) Execute(input campaign.CampaignCrate) (*campaign.Campaign, error) {
	newCampaign, err := campaign.NewCampaign(input)
	if err != nil {
		return nil, err
	}

	c, err := u.campaignRepository.Save(newCampaign)
	if err != nil {
		return nil, err
	}
	return c, nil
}
