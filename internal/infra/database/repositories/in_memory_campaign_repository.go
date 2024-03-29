package repositories

import (
	"errors"

	"github.com/bemura/campaign-management/internal/domain/campaign"
)

type inMemoryCampaignRepository struct {
	db []*campaign.Campaign
}

func NewInMemoryCampaignRepository(memoryDb []*campaign.Campaign) *inMemoryCampaignRepository {
	return &inMemoryCampaignRepository{
		db: memoryDb,
	}
}

func (r *inMemoryCampaignRepository) FindAll() ([]*campaign.Campaign, error) {
	return r.db, nil
}

func (r *inMemoryCampaignRepository) FindById(id string) (*campaign.Campaign, error) {
	var camp *campaign.Campaign
	for _, c := range r.db {
		if c.ID == id {
			camp = c
			break
		}
	}
	return camp, nil
}

func (r *inMemoryCampaignRepository) Save(camp *campaign.Campaign) (*campaign.Campaign, error) {
	for i, existingCamp := range r.db {
		if existingCamp.ID == camp.ID {
			r.db[i] = camp
			return camp, nil
		}
	}

	r.db = append(r.db, camp)
	return camp, nil
}

func (r *inMemoryCampaignRepository) Delete(id string) error {
	index := -1
	for i, c := range r.db {
		if c.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("NOT_FOUND: campaign not found")
	}

	r.db = append(r.db[:index], r.db[index+1:]...)
	return nil
}
