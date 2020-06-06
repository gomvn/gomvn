package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"
)

func (s *Server) normalizePath(path string) string {
	if path[0] == '/' {
		path = path[1:]
	}
	if strings.Contains(path, "..") || strings.Contains(path, "~") {
		return ""
	}
	if strings.Count(path, "/") <= 1 {
		return path
	}
	for _, repo := range s.repository {
		if strings.HasPrefix(path, repo) {
			return path
		}
	}
	return s.repository[0] + "/" + path
}

func (s *Server) parsePath(c *fiber.Ctx) (string, error) {
	path := s.normalizePath(c.Path())
	if strings.Count(path, "/") < 2 {
		return "", fmt.Errorf("path should be repository/group/artifact")
	}
	return path, nil
}
