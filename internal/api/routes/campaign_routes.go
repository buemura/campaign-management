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
	getAllUsecase := usecases.NewCampaignGetAllUsecase(repo)
	updateUsecase := usecases.NewCampaignUpdateUsecase(repo)
	deleteUsecase := usecases.NewCampaignDeleteUsecase(repo)
	campaingController := controllers.NewCampaignController(
		*createUsecase,
		*getAllUsecase,
		*updateUsecase,
		*deleteUsecase,
	)

	e.POST("/api/campaign", campaingController.Create)
	e.GET("/api/campaign", campaingController.Get)
	e.PUT("/api/campaign/:id", campaingController.Update)
	e.DELETE("/api/campaign/:id", campaingController.Delete)
}
