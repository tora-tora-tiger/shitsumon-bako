package database

import (
	"backend/internal/database/model"
)

func Migrate() {
	DB.AutoMigrate(&model.ImageFile{})
	DB.AutoMigrate(&model.Question{})
}