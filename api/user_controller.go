package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tuannm-sns/auth-svc/domain"
)

// ArticleHandler  represent the httphandler for article
type UserController struct {
	userUsecase domain.UserUsecase
}

func RegisterUserController(e *echo.Echo, userUsecase domain.UserUsecase) {
	controller := &UserController{
		userUsecase: userUsecase,
	}
	e.GET("/users/:id", controller.GetUser)
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
