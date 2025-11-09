package database

import (
	"fmt"

	"backend/internal/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	cfg := config.Load()
	dsn := fmt.Sprintf("%s.db", cfg.DB.Name)
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
