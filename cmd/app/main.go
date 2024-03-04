package main

import (
	// "context"
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"os"
	"os/signal"
	"time"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	"github.com/motchai-sns/sn-mono/internal/app/controller"
	"github.com/motchai-sns/sn-mono/internal/infras/connection"
	"github.com/motchai-sns/sn-mono/internal/infras/repository/models"
	"github.com/motchai-sns/sn-mono/internal/usecase"
)

func main() {

	conn := connection.NewPostgresConnection()
	// Migrate the schema
	conn.AutoMigrate(&model.User{})

	// Frameworks
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.Logger.SetLevel(log.INFO)

	// global middleware
	e.Use(middleware.Recover())
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
	// e.Pre(middleware.HTTPSRedirect())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, c.Response().Header().Get(echo.HeaderXRequestID))
	})

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})

	// register router + handler using DIContainer (wire)
	userController := InitializeUserController(conn)
	userController.RegisterHandler(e)

	authUsecase := usecase.NewAuthUsecase()
	authController := controller.NewAuthController(authUsecase)
	authController.RegisterHandler(e)

	// gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		cert, err := tls.LoadX509KeyPair("./cert/127.0.0.1.pem", "./cert/127.0.0.1-key.pem")
		if err != nil {
			fmt.Print(err.Error())
		}
		autoTLSManager := autocert.Manager{
			Prompt: autocert.AcceptTOS,
			// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
			Cache:      autocert.DirCache("/var/www/.cache"),
			HostPolicy: autocert.HostWhitelist("127.0.0.1"),
		}
		s := http.Server{
			Addr:    ":1323",
			Handler: e, // set Echo as handler
			TLSConfig: &tls.Config{
				Certificates:   []tls.Certificate{cert}, // <-- s.ListenAndServeTLS will populate this field
				GetCertificate: autoTLSManager.GetCertificate,
				NextProtos:     []string{acme.ALPNProto},
			},
			//ReadTimeout: 30 * time.Second, // use custom timeouts
		}
		// if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
		if err := s.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("An error ocured %v", err)
			e.Logger.Fatal("Shuting down server...")
		}
		// This line will only be reached if ListenAndServeTLS succeeds
	}()

	e.Logger.Info("App is listening on port :1323")
	// Wait for interrupt signal with a timeout of 30 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
