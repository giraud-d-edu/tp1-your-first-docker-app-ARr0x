package config

import (
	"animal-api/database"
	"animal-api/database/dbmodel"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	// Connexion aux repositories
	AgeEntryRepository    dbmodel.AgeEntryRepository
	AnimalSoundRepository dbmodel.AnimalSoundRepository
}

func New() (*Config, error) {
	config := Config{}

	// Initialisation de la connexion à la base de données
	databaseSession, err := gorm.Open(sqlite.Open("animal_api.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories
	config.AgeEntryRepository = dbmodel.NewAgeEntryRepository(databaseSession)
	config.AnimalSoundRepository = dbmodel.NewAnimalSoundRepository(databaseSession)

	return &config, nil
}
