package user

import "time-tracker/internal/user/ports"

type Handler struct {
	UserService ports.UserService
}
