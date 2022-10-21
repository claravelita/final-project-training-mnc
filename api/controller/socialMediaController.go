package controller

import (
	"fmt"
	"strconv"

	authLibs "github.com/claravelita/final-project-training-mnc/api/middleware"
	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"github.com/claravelita/final-project-training-mnc/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type socialMediaController struct {
	usecase usecase.SocialMediaUsecase
}

func NewSocialMediaController(usecase usecase.SocialMediaUsecase) *socialMediaController {
	return &socialMediaController{usecase: usecase}
}

func (c *socialMediaController) Route(group *echo.Group) {
	group.POST("", c.Create, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.GET("", c.GetAll, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.PUT("/:socialMediaId", c.Update, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.DELETE("/:socialMediaId", c.Delete, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
}

// Create godoc
// @Summary Create social media
// @Description This endpoint for create social media
// @Tags Social Medias
// @Accept  json
// @Produce  json
// @Param services body dtos.SocialMediaRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerCreatedResponses
// @Router /socialmedias [post]
// @Security BearerToken
func (c *socialMediaController) Create(ctx echo.Context) error {
	request := dtos.SocialMediaRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.CreateSocialMedia(convTokenID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// GetAll godoc
// @Summary Get all social medias
// @Description This endpoint for get all social medias
// @Tags Social Medias
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /socialmedias [get]
// @Security BearerToken
func (c *socialMediaController) GetAll(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	responses, err := c.usecase.GetAllSocialMedia(convTokenID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Update godoc
// @Summary Update social medias
// @Description This endpoint for update social medias
// @Tags Social Medias
// @Accept  json
// @Produce  json
// @Param socialMediaId path int true "SocialMediaID"
// @Param services body dtos.SocialMediaRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /socialmedias/{socialMediaId} [put]
// @Security BearerToken
func (c *socialMediaController) Update(ctx echo.Context) error {
	request := dtos.SocialMediaRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.UpdateSocialMedia(convTokenID, convParamID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Delete godoc
// @Summary Delete social medias
// @Description This endpoint for delete social medias
// @Tags Social Medias
// @Accept  json
// @Produce  json
// @Param socialMediaId path int true "SocialMediaID"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /socialmedias/{socialMediaId} [delete]
// @Security BearerToken
func (c *socialMediaController) Delete(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		return err
	}

	responses, err := c.usecase.DeleteSocialMedia(convTokenID, convParamID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
