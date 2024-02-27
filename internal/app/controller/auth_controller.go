package controller
//
// import (
// 	"github.com/motchai-sns/sn-mono/internal/domain"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// )
//
// type AuthController struct {
// 	userUsecase domain.IUserUsecase
// }
//
// func NewAuthController(userUsecase domain.IUserUsecase) AuthController {
// 	return AuthController{userUsecase}
// }
//
// func (uc *AuthController) RegisterHandler(e *echo.Echo) {
// 	e.GET("/oauth/google/callback", uc.oauthCallback)
// }
//
// func (uc *AuthController) oauthCallback(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
// 	}
//
// 	ctx := c.Request().Context()
// 	user, err := uc.userUsecase.GetUserByID(ctx, uint(id))
// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}
//
// 	return c.JSON(http.StatusOK, user)
// }
