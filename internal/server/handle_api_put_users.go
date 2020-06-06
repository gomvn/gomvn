package server

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
)

func (s *Server) handleApiPutUsers(c *fiber.Ctx) {
	r := new(apiPutUsersRequest)
	if err := c.BodyParser(r); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}
	if err := r.validate(); err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	user, err := s.us.Update(uint(id), r.Deploy, r.Allowed)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}

	c.JSON(&apiPutUsersResponse{
		Id:    user.ID,
		Name:  user.Name,
	})
}

type apiPutUsersRequest struct {
	Deploy  bool     `json:"deploy"`
	Allowed []string `json:"allowed"`
}

func (r *apiPutUsersRequest) validate() error {
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

type apiPutUsersResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
}
