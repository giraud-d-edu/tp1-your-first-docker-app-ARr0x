package dbmodel

import (
	"gorm.io/gorm"
)

type AgeEntry struct {
	gorm.Model
	HumanAge int `json:"human_age"`
	CatAge   int `json:"cat_age"`
}

type AgeEntryRepository interface {
	Create(entry *AgeEntry) (*AgeEntry, error)
	FindAll() ([]*AgeEntry, error)
}

type ageEntryRepository struct {
	db *gorm.DB
}

func NewAgeEntryRepository(db *gorm.DB) AgeEntryRepository {
	return &ageEntryRepository{db: db}
}

func (r *ageEntryRepository) Create(entry *AgeEntry) (*AgeEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *ageEntryRepository) FindAll() ([]*AgeEntry, error) {
	var entries []*AgeEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
