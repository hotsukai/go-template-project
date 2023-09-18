package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"sample/pkg/infrastructure/database"
	"sample/pkg/infrastructure/environment"
	middleware2 "sample/pkg/presentation/web/middleware"
)

func init() {
	environment.EnvLoad()
	database.ConnectDB()
	database.AutoMigrate()
}

func Serve() {
	defer func() {
		db, _ := database.DB.DB()
		db.Close()
	}()

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware2.SetRequestIDContext)

	e.Use(middleware.LoggerWithConfig(middleware2.CustomLogger()))

	e.HTTPErrorHandler = middleware2.CustomHTTPErrorHandler

	e.Logger.SetHeader("${time_rfc3339} ${level}")
	e.Logger.SetLevel(log.INFO)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	uh := NewUserHandler()
	e.POST("/users", func(c echo.Context) error {
		return uh.CreateUser(c)
	})
	e.GET("/users", func(c echo.Context) error {
		return uh.ListUsers(c)
	})
	e.GET("/users/:id", func(c echo.Context) error {
		return uh.ShowUser(c)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
