// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/motchai-sns/sn-mono/internal/app/controller"
	"github.com/motchai-sns/sn-mono/internal/usecase"
	"github.com/motchai-sns/sn-mono/repository/pg"
	"gorm.io/gorm"
)

// Injectors from wire.go:

// db connection will be injected by hand
func InitializeUserController(conn *gorm.DB) controller.UserController {
	iUserRepository := pg.NewPgUserRepository(conn)
	iUserUsecase := usecase.NewUserUsecase(iUserRepository)
	userController := controller.NewUserController(iUserUsecase)
	return userController
}