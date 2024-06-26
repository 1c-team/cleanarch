package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/1c-team/cleanarch/internal/domain"
	"github.com/1c-team/cleanarch/internal/infras/configs"
)

type AuthController struct {
	au domain.IAuthUsecase
}

func NewAuthController(au domain.IAuthUsecase) AuthController {
	return AuthController{au}
}

func (ac *AuthController) RegisterHandler(e *echo.Echo) {
	e.GET("/oauth/google/login", ac.login)
	e.GET("/oauth/google/callback", ac.googleCallback)
}

func (ac *AuthController) login(c echo.Context) error {
	oauthConfig := configs.GoogleOauthConfig()
	url := oauthConfig.AuthCodeURL(os.Getenv("GOOGLE_CLIENT_ID"))
	return c.JSON(200, url)
}

func (ac *AuthController) googleCallback(c echo.Context) error {
	state := c.QueryParam("state")
	if state != os.Getenv("GOOGLE_CLIENT_ID") {
		return c.JSON(400, "States don't Match!!")
	}

	code := c.QueryParam("code")
	oauthConfig := configs.GoogleOauthConfig()

	token, err := oauthConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		fmt.Printf("Error exchanging code for token: %v\n", err)
		return c.JSON(400, "Code-Token Exchange Failed")
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("Error fetching user data: %v\n", err)
		return c.JSON(400, "User Data Fetch Failed")
	}
	defer resp.Body.Close()

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing user data: %v\n", err)
		return c.JSON(400, "JSON Parsing Failed")
	}

	return c.JSON(200, string(userData))
}
