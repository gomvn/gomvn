package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"

	"github.com/gomvn/gomvn/internal/server/basicauth"
	"github.com/gomvn/gomvn/internal/service"
	"github.com/gomvn/gomvn/internal/service/user"
)

func NewRepoAuth(us *user.Service, ps *service.PathService, needsDeploy bool) func(*fiber.Ctx) {
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
			if err != nil && needsDeploy {
				log.Printf("path error: %v", err)
				return false
			}

			paths, err := us.GetPaths(u)
			if err != nil {
				log.Printf("cannot fetch user paths: %v", err)
				return false
			}

			current = "/" + current
			for _, path := range paths {
				if strings.HasPrefix(current, path.Path) && (path.Deploy || !needsDeploy) {
					return true
				}
			}

			log.Printf("not found allowed path for '%s', user paths: %v", current, paths)
			return false
		},
	})
}
