package user

import (
	"UserService/internal/middleware"
	"UserService/internal/middlewarejwt"

	"github.com/gorilla/mux"
)

func SetupRoutes(api *mux.Router, h *Handler) {
	user := api.PathPrefix("/users").Subrouter()

	// Interno (sin JWT)
	internal := user.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("/admin", h.CreateAdminInternal).Methods("POST")

	// Protegido (JWT)
	protected := user.NewRoute().Subrouter()
	protected.Use(middlewarejwt.JWTMiddleware)

	protected.HandleFunc("/{id}", h.Update).Methods("PUT")
	protected.HandleFunc("/{id}", h.Delete).Methods("DELETE")
	protected.HandleFunc("/tenant", h.ListByTenant).Methods("GET")
	protected.HandleFunc("", h.Create).Methods("POST")
}
