package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"linkmate/internal/config"
	"linkmate/internal/users"
	"net/http"
	"time"
)

type RegisterRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponseDto struct {
	UserId int `json:"userId"`
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	AccessToken string `json:"accessToken"`
}

type Controller struct {
	UsersRepository *users.Repository
	Config          *config.Config
}

func NewController(repository *users.Repository, config2 *config.Config) *Controller {
	return &Controller{UsersRepository: repository, Config: config2}
}

// RegistrationHandler godoc
//
//	@Summary	Register account
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		account	body		auth.RegisterRequestDto	true	"Add account"
//	@Success	200		{object}	RegisterResponseDto
//	@Router		/auth/register [post]
func (c Controller) RegistrationHandler(ec echo.Context) error {
	dto := new(RegisterRequestDto)

	if err := ec.Bind(dto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user, err := c.UsersRepository.CreateUser(users.CreateUserDto{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return ec.JSON(http.StatusCreated, RegisterResponseDto{
		UserId: user.Id,
	})
}

// LoginHandler godoc
//
//	@Summary	Login
//	@Router		/auth/login [post]
//	@Param		hello	body	auth.LoginRequestDto	true	"Hello"
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	auth.LoginResponseDto
func (c Controller) LoginHandler(ec echo.Context) error {
	reqDto := new(LoginRequestDto)

	if err := ec.Bind(reqDto); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user, err := c.UsersRepository.FindUserByEmailWithPassword(reqDto.Email)
	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("user with email = \"%s\" not found", reqDto.Email))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqDto.Password)); err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Password not match")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    user.Id,
			"email": user.Email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		},
	)
	result, err := token.SignedString([]byte(c.Config.JwtSecret))
	if err != nil {
		ec.Logger().Error(err)
		return echo.NewHTTPError(http.StatusNotFound, "Token create error")
	}
	return ec.JSON(http.StatusOK, LoginResponseDto{AccessToken: result})
}
