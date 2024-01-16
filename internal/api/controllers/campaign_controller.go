package controllers

import (
	"net/http"

	"github.com/bemura/campaign-management/internal/domain/campaign"
	"github.com/bemura/campaign-management/internal/usecases"
	"github.com/labstack/echo/v4"
)

type CampaignController struct {
	campaignCreateUsecase usecases.CampaignCreateUsecase
}

func NewCampaignController(campaignCreateUsecase usecases.CampaignCreateUsecase) *CampaignController {
	return &CampaignController{
		campaignCreateUsecase: campaignCreateUsecase,
	}
}

func (cc *CampaignController) Create(c echo.Context) error {
	input := new(campaign.CampaignCrate)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Unable to parse request body",
		})
	}

	res, err := cc.campaignCreateUsecase.Execute(*input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}