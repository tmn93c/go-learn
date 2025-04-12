//go:build wireinject
// +build wireinject

package di

import (
	"basic-to-advanced/di/google-wire-di/logger"
	"basic-to-advanced/di/google-wire-di/repo"
	"basic-to-advanced/di/google-wire-di/service"

	"github.com/google/wire"
)

func InitializeUserService() *service.UserService {
	wire.Build(
		logger.NewLogger,
		repo.NewMockUserRepo,
		service.NewUserService,
	)
	return nil
}
