package main

import (
	"github.com/bemura/campaign-management/config"
	"github.com/bemura/campaign-management/internal/api/routes"
	"github.com/labstack/echo/v4"
)

func init() {
	config.LoadConfigs()
}

func main() {
	e := echo.New()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":" + config.PORT))
}
