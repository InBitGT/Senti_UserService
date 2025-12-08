package routes

import (
	"UserService/internal/modules/user"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetUserHandler() *user.Handler
}

func SetupRoutes(r *mux.Router, handlers RouteHandlers) {
	api := r.PathPrefix("/api/v1").Subrouter()

	user.SetupRoutes(api, handlers.GetUserHandler())
}
