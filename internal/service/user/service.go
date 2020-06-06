package user

import (
	"github.com/jinzhu/gorm"
)

const (
	TokenLength = 36
	BcryptCost = 12
)

func New(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

type Service struct {
	db *gorm.DB
}
