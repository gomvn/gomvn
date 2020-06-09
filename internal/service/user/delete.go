package user

import (
	"github.com/jinzhu/gorm"

	"github.com/gomvn/gomvn/internal/entity"
)

func (s *Service) Delete(id uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.db.Where("user_id = ?", id).Delete(entity.Path{}).Error; err != nil {
			return err
		}
		if err := s.db.Where("id = ?", id).Delete(entity.User{}).Error; err != nil {
			return err
		}

		return nil
	})
}
