package routes

import (
	"github.com/bemura/campaign-management/internal/api/controllers"
	"github.com/bemura/campaign-management/internal/domain/campaign"
	"github.com/bemura/campaign-management/internal/infra/database/repositories"
	"github.com/bemura/campaign-management/internal/usecases"
	"github.com/labstack/echo/v4"
)

func setupCampaignRoutes(e *echo.Echo) {
	var camp []*campaign.Campaign
	repo := repositories.NewInMemoryCampaignRepository(camp)
	createUsecase := usecases.NewCampaignCreateUsecase(repo)
	campaingController := controllers.NewCampaignController(*createUsecase)

	e.GET("/campaign", campaingController.Create)
	e.POST("/campaign", campaingController.Create)
	e.PUT("/campaign/:id", campaingController.Create)
	e.DELETE("/campaign/:id", campaingController.Create)
}
