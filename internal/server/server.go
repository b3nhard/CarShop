package server

import (
	"net/http"
	"os"
	"path/filepath"

	mw "github.com/b3nhard/car-rental/internal/middleware"
	"github.com/b3nhard/car-rental/internal/router"
	"github.com/b3nhard/car-rental/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type Server struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	r := chi.NewRouter()

	return &Server{
		Router: r,
		DB:     db,
	}
}

func (s *Server) MountRoutes() {
	// Middewares
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Compress(5))
	s.Router.Use(mw.HTMXMiddleware)

	// Serve from Static, Public and Upload Directories
	workDir, _ := os.Getwd()
	publicDir := http.Dir(filepath.Join(workDir, "web/public"))
	staticDir := http.Dir(filepath.Join(workDir, "web/static"))
	uploadDir := http.Dir(filepath.Join(workDir, "web/public"))
	utils.FileServer(s.Router, "/", publicDir)
	utils.FileServer(s.Router, "/static", staticDir)
	utils.FileServer(s.Router, "/upload", uploadDir)

	// Router routes
	router.SetupRoutes(s.Router)

}
