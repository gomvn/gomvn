package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/gomvn/gomvn/internal/entity"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{})

	return db, nil
}
