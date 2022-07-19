package migrations

import (
	"github.com/albadauto/projeto/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Books{})
}
