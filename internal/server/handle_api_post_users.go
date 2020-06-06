package server

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func (s *Server) handleApiPostUsers(c *fiber.Ctx) {
	r := new(apiPostUsersRequest)
	if err := c.BodyParser(r); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}
	if err := r.validate(); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	user, token, err := s.us.Create(r.Name, r.Admin, r.Deploy, r.Allowed)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}

	c.JSON(&apiPostUsersResponse{
		Id:    user.ID,
		Name:  user.Name,
		Token: token,
	})
}

type apiPostUsersRequest struct {
	Name    string   `json:"name"`
	Admin   bool     `json:"admin"`
	Deploy  bool     `json:"deploy"`
	Allowed []string `json:"allowed"`
}

func (r *apiPostUsersRequest) validate() error {
	if r.Name == "" {
		return fmt.Errorf("field 'name' cannot be empty")
	}
	if len(r.Allowed) < 1 {
		return fmt.Errorf("field 'allowed' must contain at least one string")
	}
	for _, path := range r.Allowed {
		if path[0] != '/' {
			return fmt.Errorf("paths in field 'allowed' must start with '/'")
		}
	}
	return nil
}

type apiPostUsersResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
