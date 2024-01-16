package usecases

import "github.com/bemura/campaign-management/internal/domain/campaign"

type CampaignUpdateUsecase struct {
	campaignRepository campaign.CampaignRepository
}

func NewCampaignUpdateUsecase(campaignRepository campaign.CampaignRepository) *CampaignUpdateUsecase {
	return &CampaignUpdateUsecase{
		campaignRepository: campaignRepository,
	}
}

func (u *CampaignUpdateUsecase) Execute(input campaign.CampaignUpdate) (*campaign.Campaign, error) {
	c, err := u.campaignRepository.FindById(input.ID)
	if err != nil {
		return nil, err
	}

	if updateCampaignFields(c, input) {
		savedCampaign, err := u.campaignRepository.Save(c)
		if err != nil {
			return nil, err
		}
		return savedCampaign, nil
	}

	return c, nil
}

func updateCampaignFields(c *campaign.Campaign, input campaign.CampaignUpdate) bool {
	changed := false

	if input.Name != nil && *input.Name != c.Name {
		c.Name = *input.Name
		changed = true
	}
	if input.Status != nil && *input.Status != c.Status {
		c.Status = *input.Status
		changed = true
	}

	return changed
}
