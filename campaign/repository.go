package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userId int) ([]Campaign, error)
	FindById(Id int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary= 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserId(userId int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindById(Id int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Where("id = ?", Id).Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}