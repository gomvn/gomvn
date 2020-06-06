package middleware

import (
	"strings"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"github.com/gomvn/gomvn/internal/server/basicauth"
	"github.com/gomvn/gomvn/internal/service"
	"github.com/gomvn/gomvn/internal/service/user"
)

func NewPutAuth(us *user.Service, ps *service.PathService) func(*fiber.Ctx) {
	return basicauth.New(basicauth.Config{
		Authorizer: func(c *fiber.Ctx, name string, token string) bool {
			u, err := us.GetByName(name)
			if err != nil {
				return false
			}
			if err := bcrypt.CompareHashAndPassword([]byte(u.TokenHash), []byte(token)); err != nil {
				return false
			}

			_, current, _, err := ps.ParsePathParts(c)
			if err != nil {
				return false
			}

			paths, err := us.GetPaths(u)
			if err != nil {
				return false
			}

			current = "/" + current
			for _, path := range paths {
				if strings.HasPrefix(current, path.Path) && path.Deploy {
					return true
				}
			}

			return false
		},
	})
}
