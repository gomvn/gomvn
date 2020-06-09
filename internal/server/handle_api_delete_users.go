package server

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func (s *Server) handleApiDeleteUsers(c *fiber.Ctx) {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	if err := s.us.Delete(uint(id)); err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		return
	}
}
