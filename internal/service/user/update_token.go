package user

import (
	"time"

	"github.com/jinzhu/gorm"
	gonanoid "github.com/matoous/go-nanoid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gomvn/gomvn/internal/entity"
)

func (s *Service) UpdateToken(id uint) (*entity.User, string, error) {
	var user entity.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, "", err
	}

	token, err := gonanoid.ID(TokenLength)
	if err != nil {
		return nil, "", err
	}

	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), BcryptCost)
	if err != nil {
		return nil, "", err
	}


	err = s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&user).
			Updates(map[string]interface{}{
				"TokenHash": tokenHash,
				"UpdatedAt": time.Now(),
			}).
			Error
	})
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}
