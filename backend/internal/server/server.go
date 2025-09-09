package server

import (
	"net/http"

	"backend/internal/config"
	"backend/internal/handler"
	"backend/internal/middleware"
	"backend/pkg/schema"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.Config
}

func New(cfg *config.Config) *Server {
	e := echo.New()
	
	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(middleware.CORS())

	return &Server{
		echo:   e,
		config: cfg,
	}
}

func (s *Server) SetupRoutes() {
	// Health check endpoint
	s.echo.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "質問箱API Server is running"})
	})

	// Register API routes with handlers
	apiHandler := handler.NewAPIHandler()
	schema.RegisterHandlers(s.echo, apiHandler)
}

func (s *Server) Start() error {
	s.SetupRoutes()
	return s.echo.Start(":" + s.config.Server.Port)
}