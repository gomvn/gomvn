package user

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/gomvn/gomvn/internal/entity"
)

func (s *Service) Update(id uint, deploy bool, paths []string) (*entity.User, error) {
	var user entity.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		for _, path := range paths {
			userPath := entity.Path{
				UserID: user.ID,
				Path:   path,
			}
			q := tx.Assign(map[string]interface{}{"Deploy": deploy, "UpdatedAt": time.Now()}).
				FirstOrCreate(&userPath)
			if err := q.Error; err != nil {
				return err
			}
		}

		return tx.Model(&user).
			Updates(map[string]interface{}{"UpdatedAt": time.Now()}).
			Error
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
