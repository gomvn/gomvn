package server

import (
	"strconv"

	"github.com/gofiber/fiber"
)

func getQueryUint64(c *fiber.Ctx, name string, def uint64) uint64 {
	if val, err := strconv.ParseUint(c.Query(name), 10, 64); err == nil {
		return val
	}
	return def
}
