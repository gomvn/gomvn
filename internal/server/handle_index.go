package server

import (
	"github.com/gofiber/fiber"
)

func (s *Server) handleIndex(c *fiber.Ctx) {
	c.Render("index", fiber.Map{
		"Name": s.name,
	})
}
