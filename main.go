package main

import (
	"database/sql"
	"errors"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
	"linkmate/internal/auth"
	"linkmate/internal/config"
	"linkmate/internal/links"
	"linkmate/internal/users"
	"net/http"
)

//	@title		Linkmate API
//	@version	1.0

//	@contact.name	Sam Bulatov
//	@contact.url	https://mephi.dev
//	@contact.email	sam@mephi.dev
//	@host			localhost:9000
//	@BasePath		/api

// @license.name				MIT
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	conf := config.New()
	app := echo.New()

	db, err := sql.Open("postgres", conf.DatabaseConnectUrl)

	if err != nil {
		app.Logger.Fatal(errors.New("database connect error"))
	}

	err = db.Ping()

	if err != nil {
		app.Logger.Fatal(errors.New("database connect error"))
	}

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("X-Powered-by", "Linkmate")
			return next(c)
		}
	})

	usersRepo := users.NewRepository(db)
	linksRepo := links.NewRepository(db)

	authController := auth.NewController(usersRepo, conf)
	linkController := links.NewController(linksRepo, conf)

	authGroup := app.Group("/api/auth")
	internalGroup := app.Group("/api", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(conf.JwtSecret),
	}))

	app.GET(":key", linkController.RedirectLinkHandler)
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	authGroup.POST("/register", authController.RegistrationHandler)
	authGroup.POST("/login", authController.LoginHandler)
	internalGroup.POST("/links", linkController.CreateLinkHandler)
	internalGroup.GET("/links", linkController.LinkListHandler)

	app.Logger.Fatal(app.Start(fmt.Sprintf(":%d", conf.Port)))
}
