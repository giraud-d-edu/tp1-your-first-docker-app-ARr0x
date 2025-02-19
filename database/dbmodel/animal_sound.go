package dbmodel

import (
	"gorm.io/gorm"
)

// Modèle AnimalSound
type AnimalSound struct {
	gorm.Model
	AnimalName string `json:"animal"`
	Sound      string `json:"sound"`
}

// Interface AnimalSoundRepository
type AnimalSoundRepository interface {
	Create(entry *AnimalSound) (*AnimalSound, error)
	FindAll() ([]*AnimalSound, error)
}

// Struct qui implémente AnimalSoundRepository
type animalSoundRepository struct {
	db *gorm.DB
}

// Constructeur pour AnimalSoundRepository
func NewAnimalSoundRepository(db *gorm.DB) AnimalSoundRepository {
	return &animalSoundRepository{db: db}
}

// Méthode Create pour insérer un AnimalSound dans la BDD
func (r *animalSoundRepository) Create(entry *AnimalSound) (*AnimalSound, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

// Méthode FindAll pour récupérer tous les enregistrements AnimalSound
func (r *animalSoundRepository) FindAll() ([]*AnimalSound, error) {
	var entries []*AnimalSound
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
