package main

import (
	"log/slog"
	"os"

	"github.com/bemura/campaign-management/config"
	"github.com/bemura/campaign-management/internal/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadConfigs()
}

func main() {
	slogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(slogger)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":" + config.PORT))
}
