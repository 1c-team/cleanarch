package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/motchai-sns/sn-mono/connection"
	"github.com/motchai-sns/sn-mono/repository/models"
)

func main() {
	conn := connection.NewPostgresConnection()
	// Migrate the schema
	conn.AutoMigrate(&model.User{})

	// Frameworks
	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	// global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.CSRF())
	e.Use(middleware.CORS())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Printf(c.Path())
		},
		Timeout: 60 * time.Second,
	}))
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Response().Header().Get(echo.HeaderXRequestID))
	})

	// register router + handler using DIContainer (wire)
	userController := InitializeUserController(conn)
	userController.RegisterHandler(e)

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
