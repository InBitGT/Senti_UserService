package server

import (
	"UserService/internal/modules/user"

	"gorm.io/gorm"
)

func InitializeHandlers(db *gorm.DB) *Handlers {

	// USER MODULE
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	return &Handlers{
		User: userHandler,
	}
}
