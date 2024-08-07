package services

import "time-tracker/internal/user/ports"

type Service struct {
	Repo ports.UserRepository
}
