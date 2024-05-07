//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/1c-team/cleanarch/internal/app/controller"
	"github.com/1c-team/cleanarch/internal/usecase"
	"github.com/1c-team/cleanarch/internal/infras/repository/pg"
)

// db connection will be injected by hand
func InitializeUserController(conn *gorm.DB) controller.UserController {
	wire.Build(
		pg.NewPgUserRepository,
		usecase.NewUserUsecase,
		controller.NewUserController,
	)
	return controller.UserController{}
}
