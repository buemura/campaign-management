package usecases

import "github.com/bemura/campaign-management/internal/domain/campaign"

type CampaignDeleteUsecase struct {
	campaignRepository campaign.CampaignRepository
}

func NewCampaignDeleteUsecase(campaignRepository campaign.CampaignRepository) *CampaignDeleteUsecase {
	return &CampaignDeleteUsecase{
		campaignRepository: campaignRepository,
	}
}

func (u *CampaignDeleteUsecase) Execute(id string) error {
	err := u.campaignRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
