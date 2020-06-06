package user

import (
	"github.com/jinzhu/gorm"
	"github.com/matoous/go-nanoid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gomvn/gomvn/internal/entity"
)

func (s *Service) Create(name string, api bool, deploy bool, paths []string) (*entity.User, string, error) {
	token, err := gonanoid.ID(TokenLength)
	if err != nil {
		return nil, "", err
	}

	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), BcryptCost)
	if err != nil {
		return nil, "", err
	}

	user := entity.User{
		Name:      name,
		Admin:     api,
		TokenHash: string(tokenHash),
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		for _, path := range paths {
			userPath := entity.Path{
				UserID: user.ID,
				Path:   path,
				Deploy: deploy,
			}
			if err := tx.Create(&userPath).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}
