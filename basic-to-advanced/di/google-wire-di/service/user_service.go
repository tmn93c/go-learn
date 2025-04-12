package service

import (
	"basic-to-advanced/di/google-wire-di/logger"
	"basic-to-advanced/di/google-wire-di/repo"
)

type UserService struct {
	Repo   repo.UserRepo
	logger *logger.Logger
}

func NewUserService(r repo.UserRepo, l *logger.Logger) *UserService {
	return &UserService{Repo: r, logger: l}
}

func (s *UserService) GetUserName(id int) string {
	user := s.Repo.GetUser(id)
	s.logger.Info("Fetched user: " + user.Name)
	if user == nil {
		return "Not found"
	}
	return user.Name
}
