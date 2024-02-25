//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/tuannm-sns/auth-svc/api"
	"github.com/tuannm-sns/auth-svc/internal/repository/pg"
	"github.com/tuannm-sns/auth-svc/internal/usecase"
)

// db connection will be injected by hand
func InitializeUserController(conn *gorm.DB) api.UserController {
	wire.Build(
		pg.NewPgUserRepository,
		usecase.NewUserUsecase,
		// wire.Bind(new(domain.UserRepository), new(*pg.PgUserRepository)),
		// wire.Bind(new(domain.UserUsecase), new(*usecase.UserUsecase)),
		api.NewUserController,
	)
	return api.UserController{}
}
