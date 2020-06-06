package user

import (
	"github.com/jinzhu/gorm"

	"github.com/gomvn/gomvn/internal/entity"
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

func (s *Service) GetByName(name string) (*entity.User, error) {
	var user entity.User
	if err := s.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
