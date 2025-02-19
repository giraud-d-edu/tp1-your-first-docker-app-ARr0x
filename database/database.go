package database

import (
	"animal-api/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.AgeEntry{},
		&dbmodel.AnimalSound{},
	)
	log.Println("Database migrated successfully")
}
