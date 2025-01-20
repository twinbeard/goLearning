//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/twinbeard/goLearning/internal/controller"
	"github.com/twinbeard/goLearning/internal/repo"
	"github.com/twinbeard/goLearning/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil

}
