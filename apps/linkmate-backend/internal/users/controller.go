package users

import (
	"github.com/labstack/echo/v4"
	"linkmate/internal/config"
	authUtils "linkmate/internal/shared/auth-utils"
	"net/http"
)

type Controller struct {
	config          *config.Config
	usersRepository *Repository
}

func NewController(repository *Repository, config2 *config.Config) *Controller {
	return &Controller{usersRepository: repository, config: config2}
}

// SelfHandler godoc
//
//	@Summary	Return current user
//	@Router		/users/self [get]
//	@Success	200	{object}	users.WhoAmIDto
//	@Tags		users
//	@Accept		json
//	@Produce	json
//
//	@Security	ApiKeyAuth
func (c Controller) SelfHandler(ec echo.Context) error {
	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))
	user, err := c.usersRepository.FindUserById(tokenData.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.JSON(http.StatusOK, WhoAmIDto{
		Id:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	})
}

// DeleteSelfHandler godoc
//
//	@Summary	Delete self user
//	@Router		/users/self [delete]
//	@Success	200
//	@Tags		users
//
//	@Security	ApiKeyAuth
func (c Controller) DeleteSelfHandler(ec echo.Context) error {
	tokenData, err := authUtils.ParseJwtData(ec.Get("user"))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = c.usersRepository.DeleteUserById(tokenData.UserId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.NoContent(http.StatusOK)
}
