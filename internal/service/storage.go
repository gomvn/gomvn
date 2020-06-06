package service

import (
	"bufio"
	"os"
	"strings"

	"github.com/gofiber/fiber"
)

func NewStorage() *Storage {
	return &Storage{
		root: "data/repository",
	}
}

type Storage struct {
	root string
}

func (s *Storage) GetRoot() string {
	return s.root
}

func (s *Storage) File(path string) string {
	return s.root + "/" + path
}

func (s *Storage) FileExists(path string) bool {
	file := s.File(path)
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (s *Storage) WriteFromRequest(c *fiber.Ctx, path string) error {
	file := s.File(path)
	fdir := dir(file)
	if err := os.MkdirAll(fdir, os.ModeDir); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if err := c.Fasthttp.Request.BodyWriteTo(w); err != nil {
		return err
	}
	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}

func dir(path string) string {
	index := strings.LastIndex(path, "/")
	return path[:index]
}
