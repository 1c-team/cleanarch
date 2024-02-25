package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tuannm-sns/auth-svc/internal/domain"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) UserController {
	return UserController{userUsecase}
}

func RegisterUserController(e *echo.Echo, userController UserController) {
	e.GET("/users/:id", userController.GetUser)
}

func (userController *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()
	usr, err := userController.userUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, usr)
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
