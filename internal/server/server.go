package server

import (
	"log"
	"time"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/template/html"

	"github.com/gomvn/gomvn/internal/config"
	"github.com/gomvn/gomvn/internal/service"
)

func New(conf *config.App, storage *service.Storage) *Server {
	app := fiber.New()
	app.Settings.IdleTimeout = time.Second * 5
	app.Settings.DisableStartupMessage = true
	app.Settings.Templates = html.New("./views", ".html")

	app.Use(compression.New())
	app.Use(logger.New())

	server := &Server{
		app:        app,
		name:       conf.Name,
		storage:    storage,
		listen:     conf.Server.GetListenAddr(),
		repository: conf.Repository,
	}

	app.Put("/*", server.handlePut)
	app.Get("/", server.handleIndex)
	app.Static("/", storage.GetRoot(), fiber.Static{
		Browse: true,
	})

	return server
}

type Server struct {
	app        *fiber.App
	name       string
	storage    *service.Storage
	listen     string
	repository []string
}

func (s *Server) Listen() error {
	log.Printf("GoMVN self-hosted repository listening on %s\n", s.listen)
	go s.app.Listen(s.listen)
	return nil
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
