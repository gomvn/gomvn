package middleware

import (
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"github.com/gomvn/gomvn/internal/server/basicauth"
	"github.com/gomvn/gomvn/internal/service/user"
)

func NewApiAuth(us *user.Service) func(*fiber.Ctx) {
	return basicauth.New(basicauth.Config{
		Authorizer: func(c *fiber.Ctx, name string, token string) bool {
			u, err := us.GetByName(name)
			if err != nil || !u.Admin {
				return false
			}
			if err := bcrypt.CompareHashAndPassword([]byte(u.TokenHash), []byte(token)); err != nil {
				return false
			}
			return true
		},
	})
}
