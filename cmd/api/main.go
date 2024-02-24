package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tuannm99/social-network/backend/auth/api"
	"github.com/tuannm99/social-network/backend/auth/connection"
	"github.com/tuannm99/social-network/backend/auth/internal/model"
	"github.com/tuannm99/social-network/backend/auth/internal/repository/pg"
	"github.com/tuannm99/social-network/backend/auth/internal/usecase"
)

func main() {
	conn := connection.CreatePostgresConnection()
	// Migrate the schema
	conn.AutoMigrate(&model.User{})

	// DI
	userRepo := pg.NewPgArticleRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepo, 10*time.Second)

	// Frameworks
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// register router + handler
	api.RegisterUserController(e, userUsecase)

	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("An error ocured %v", err)
			e.Logger.Fatal("Shuting down server...")
		}
	}()

	// Wait for interrupt signal with a timeout of 30 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)

	}
}
