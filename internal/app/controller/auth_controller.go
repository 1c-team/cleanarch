package controller

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/motchai-sns/sn-mono/configs"
	"github.com/motchai-sns/sn-mono/internal/domain"
)

type AuthController struct {
	authUsecase domain.IAuthUsecase
}

func NewAuthController(authUsecase domain.IAuthUsecase) AuthController {
	return AuthController{authUsecase}
}

func (authController *AuthController) RegisterHandler(e *echo.Echo) {
	e.GET("/oauth/google/login", authController.login)
	e.POST("/oauth/google/callback", authController.googleCallback)
}

func (authController *AuthController) login(c echo.Context) error {
	oauthConfig := configs.GoogleOauthConfig()
	// authController.authUsecase.Login()
	url := oauthConfig.AuthCodeURL("randomstate")
	return c.JSON(200, url)
}

func (authController *AuthController) googleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != "randomstate" {
		return c.JSON(400, "States don't Match!!")
	}

	code := c.QueryParam("code")
	oauthConfig := configs.GoogleOauthConfig()

	token, err := oauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(400, "Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return c.JSON(400, "User Data Fetch Failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(400, "JSON Parsing Failed")
	}

	return c.JSON(200, string(userData))
}
