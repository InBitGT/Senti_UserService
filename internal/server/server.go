package server

import (
	"net/http"

	"UserService/internal/middleware"
	"UserService/internal/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	db     *gorm.DB
}

func NewServer(db *gorm.DB, handlers *Handlers) *Server {

	s := &Server{
		Router: mux.NewRouter(),
		db:     db,
	}

	// MIDDLEWARES GLOBALES
	s.Router.Use(middleware.CORS)     // CORS
	s.Router.Use(middleware.Logger)   // Logging
	s.Router.Use(middleware.Recovery) // Captura panics

	// Rutas
	routes.SetupRoutes(s.Router, handlers)

	// Manejo de preflight
	s.Router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return s
}
