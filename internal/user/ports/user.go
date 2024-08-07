package ports

import "time-tracker/internal/user/domain"

type UserService interface {
	Create(user domain.User) error
}

type UserRepository interface {
	Save(user domain.User) error
}
