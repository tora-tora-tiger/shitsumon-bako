package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/internal/config"
	"backend/internal/database/query"
)

var DB *gorm.DB
var err error


func init() {
	cfg := config.Load()
	dsn := fmt.Sprintf("%s.db", cfg.DB.Name)
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	query.SetDefault(DB)

	if err != nil {
		panic("failed to connect database")
	}
}
