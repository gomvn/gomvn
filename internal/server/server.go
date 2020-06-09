package server

import (
	"log"
	"time"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/template/html"

	"github.com/gomvn/gomvn/internal/config"
	"github.com/gomvn/gomvn/internal/server/middleware"
	"github.com/gomvn/gomvn/internal/service"
	"github.com/gomvn/gomvn/internal/service/user"
)

func New(conf *config.App, ps *service.PathService, storage *service.Storage, us *user.Service, rs *service.RepoService) *Server {
	app := fiber.New()
	app.Settings.IdleTimeout = time.Second * 5
	app.Settings.DisableStartupMessage = true
	app.Settings.Templates = html.New("./views", ".html")

	app.Use(compression.New())
	app.Use(logger.New())

	server := &Server{
		app:     app,
		name:    conf.Name,
		listen:  conf.Server.GetListenAddr(),
		ps:      ps,
		storage: storage,
		us:      us,
		rs:      rs,
	}

	api := app.Group("/api")
	api.Use(middleware.NewApiAuth(us))
	api.Get("/users", server.handleApiGetUsers)
	api.Post("/users", server.handleApiPostUsers)
	api.Put("/users/:id", server.handleApiPutUsers)
	api.Delete("/users/:id", server.handleApiDeleteUsers)
	api.Get("/users/:id/refresh", server.handleApiGetUsersRefresh)

	app.Put("/*", middleware.NewRepoAuth(us, ps, true), server.handlePut)
	app.Get("/", server.handleIndex)

	app.Use(middleware.NewRepoAuth(us, ps, false))
	app.Static("/", storage.GetRoot(), fiber.Static{
		Browse: true,
	})

	return server
}

type Server struct {
	app     *fiber.App
	name    string
	listen  string
	ps      *service.PathService
	storage *service.Storage
	us      *user.Service
	rs      *service.RepoService
}

func (s *Server) Listen() error {
	log.Printf("GoMVN self-hosted repository listening on %s\n", s.listen)
	go s.app.Listen(s.listen)
	return nil
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
