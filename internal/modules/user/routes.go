package user

import (
	"UserService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, h *Handler) {

	user := r.PathPrefix("/users").Subrouter()
	user.Use(middleware.JWTMiddleware)

	user.HandleFunc("", h.Create).Methods("POST")
	user.HandleFunc("/{id}", h.Update).Methods("PUT")
	user.HandleFunc("/{id}", h.Delete).Methods("DELETE")
	user.HandleFunc("/tenant", h.ListByTenant).Methods("GET")
}
