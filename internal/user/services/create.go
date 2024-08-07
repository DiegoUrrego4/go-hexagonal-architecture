package services

import (
	"time-tracker/internal/user/domain"
)

func (s Service) Create(user domain.User) error {
	err := s.Repo.Save(user)
	if err != nil {
		return err
	}
	return nil
}
