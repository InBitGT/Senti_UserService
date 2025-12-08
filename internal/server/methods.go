package server

import "UserService/internal/modules/user"

func (h *Handlers) GetUserHandler() *user.Handler {
	return h.User
}
