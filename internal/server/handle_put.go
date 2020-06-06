package server

import (
	"log"

	"github.com/gofiber/fiber"
)

func (s *Server) handlePut(c *fiber.Ctx) {
	path, err := s.ps.ParsePath(c)
	if err != nil {
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	if err := s.storage.WriteFromRequest(c, path); err != nil {
		log.Printf("cannot put data: %v", err)
		c.Status(fiber.StatusBadRequest).SendString(err.Error())
		return
	}

	c.Status(fiber.StatusOK)
}
