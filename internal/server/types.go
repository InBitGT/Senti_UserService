package server

import "UserService/internal/modules/user"

type Handlers struct {
	User *user.Handler
}
