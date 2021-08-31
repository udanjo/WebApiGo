package migrations

import (
	"github.com/ueverson/WebApiGo/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {

	db.AutoMigrate(models.Movie{})
}
