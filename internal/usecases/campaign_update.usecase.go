package usecases

import (
	"errors"

	"github.com/bemura/campaign-management/internal/domain/campaign"
)

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
	if c == nil {
		return nil, errors.New("NOT_FOUND: campaign not found")
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

	if input.Name != nil {
		c.Name = *input.Name
		changed = true
	}

	if input.Status != nil {
		c.Status = *input.Status
		changed = true
	}

	return changed
}
