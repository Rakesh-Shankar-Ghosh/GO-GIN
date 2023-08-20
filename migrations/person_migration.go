package migrations

import (
	"backend/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&models.Person{})
}
