//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/tuannm-sns/auth-svc/internal/app/controller"
	"github.com/tuannm-sns/auth-svc/internal/usecase"
	"github.com/tuannm-sns/auth-svc/repository/pg"
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
