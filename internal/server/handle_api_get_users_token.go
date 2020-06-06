package server

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func (s *Server) handleApiGetUsersToken(c *fiber.Ctx) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	user, token, err := s.us.UpdateToken(uint(id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}

	c.JSON(&apiGetUsersTokenResponse{
		Id:    user.ID,
		Name:  user.Name,
		Token: token,
	})
}

type apiGetUsersTokenResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
