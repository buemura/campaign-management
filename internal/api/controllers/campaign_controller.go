package controllers

import (
	"log/slog"
	"net/http"

	"github.com/bemura/campaign-management/internal/api/helpers"
	"github.com/bemura/campaign-management/internal/domain/campaign"
	"github.com/bemura/campaign-management/internal/usecases"
	"github.com/labstack/echo/v4"
)

type CampaignController struct {
	campaignCreateUsecase usecases.CampaignCreateUsecase
	campaignGetAllUsecase usecases.CampaignGetAllUsecase
	campaignUpdateUsecase usecases.CampaignUpdateUsecase
	campaignDeleteUsecase usecases.CampaignDeleteUsecase
}

func NewCampaignController(
	campaignCreateUsecase usecases.CampaignCreateUsecase,
	campaignGetAllUsecase usecases.CampaignGetAllUsecase,
	campaignUpdateUsecase usecases.CampaignUpdateUsecase,
	campaignDeleteUsecase usecases.CampaignDeleteUsecase,
) *CampaignController {
	return &CampaignController{
		campaignCreateUsecase: campaignCreateUsecase,
		campaignGetAllUsecase: campaignGetAllUsecase,
		campaignUpdateUsecase: campaignUpdateUsecase,
		campaignDeleteUsecase: campaignDeleteUsecase,
	}
}

func (cc *CampaignController) Create(c echo.Context) error {
	slog.Info("[CampaignController][Create] - Creating campaign")
	input := new(campaign.CampaignCrate)
	if err := c.Bind(input); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	res, err := cc.campaignCreateUsecase.Execute(*input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}

func (cc *CampaignController) Get(c echo.Context) error {
	slog.Info("[CampaignController][Get] - Getting campaigns")
	res, err := cc.campaignGetAllUsecase.Execute()
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (cc *CampaignController) Update(c echo.Context) error {
	slog.Info("[CampaignController][Update] - Updating campaign")
	input := new(campaign.CampaignUpdate)
	if err := c.Bind(input); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}

	input.ID = c.Param("id")
	res, err := cc.campaignUpdateUsecase.Execute(*input)
	if err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (cc *CampaignController) Delete(c echo.Context) error {
	slog.Info("[CampaignController][Delete] - Deleting campaign")
	id := c.Param("id")
	if err := cc.campaignDeleteUsecase.Execute(id); err != nil {
		return helpers.BuildErrorResponse(c, err.Error())
	}
	return c.JSON(http.StatusOK, nil)
}
