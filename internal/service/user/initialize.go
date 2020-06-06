package user

import (
	"log"

	"github.com/jinzhu/gorm"

	"github.com/gomvn/gomvn/internal/entity"
)

func Initialize(db *gorm.DB, us *Service) error {
	var count int
	db.Model(&entity.User{}).Count(&count)

	if count == 0 {
		log.Println("Initializing first user...")

		user, token, err := us.Create("admin", true, true, []string{"/"})
		if err != nil {
			return err
		}

		log.Printf("USER ID: %d\n", user.ID)
		log.Printf("USERNAME: %s\n", user.Name)
		log.Printf("TOKEN: %s\n", token)
	}

	return nil
}
