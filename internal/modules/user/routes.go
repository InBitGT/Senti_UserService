package user

import (
	"UserService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, h *Handler) {
	user := r.PathPrefix("/users").Subrouter()

	protected := user.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/{id}", h.Update).Methods("PUT")
	protected.HandleFunc("/{id}", h.Delete).Methods("DELETE")
	protected.HandleFunc("/tenant", h.ListByTenant).Methods("GET")

	protected.HandleFunc("", h.Create).Methods("POST")

}
