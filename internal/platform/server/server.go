package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/m4t1t0/GoCoinTracker/internal/asset"
	"github.com/m4t1t0/GoCoinTracker/internal/platform/server/handler/createAsset"
	"github.com/m4t1t0/GoCoinTracker/internal/platform/server/handler/home"
	"log"
)

type Server struct {
	port   uint
	app    *fiber.App
	assets asset.Service
}

func New(port uint, assetsSvc asset.Service) Server {
	srv := Server{
		app:    fiber.New(),
		port:   port,
		assets: assetsSvc,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on port", s.port)
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}

func (s *Server) registerRoutes() {
	s.app.Get("/", home.Handler())
	s.app.Post("/api/v1/assets", createAsset.Handler(s.assets))
}
