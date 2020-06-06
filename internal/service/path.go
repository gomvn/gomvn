package service

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber"

	"github.com/gomvn/gomvn/internal/config"
)

func NewPathService(conf *config.App) *PathService {
	return &PathService{
		repository: conf.Repository,
	}
}

type PathService struct {
	repository []string
}

func (s *PathService) NormalizePath(path string) string {
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

func (s *PathService) ParsePath(c *fiber.Ctx) (string, error) {
	path := s.NormalizePath(c.Path())
	if strings.Count(path, "/") < 2 {
		return "", fmt.Errorf("path should be repository/group/artifact")
	}
	return path, nil
}

func (s *PathService) ParsePathParts(c *fiber.Ctx) (string, string, string, error) {
	path, err := s.ParsePath(c)
	if err != nil {
		return "", "", "", err
	}
	parts := strings.Split(path,"/")
	last := len(parts) - 1
	return parts[0], strings.Join(parts[1:last-1], "/"), parts[last], nil
}
