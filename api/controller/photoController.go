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

type photoController struct {
	usecase usecase.PhotoUsecase
}

func NewPhotoController(usecase usecase.PhotoUsecase) *photoController {
	return &photoController{usecase: usecase}
}

func (c *photoController) Route(group *echo.Group) {
	group.POST("", c.Create, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.GET("", c.GetAll, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.PUT("/:photoId", c.Update, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.DELETE("/:photoId", c.Delete, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
}

// Create godoc
// @Summary Create photo
// @Description This endpoint for create photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param services body dtos.PhotoRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerCreatedResponses
// @Router /photos [post]
// @Security BearerToken
func (c *photoController) Create(ctx echo.Context) error {
	request := dtos.PhotoRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.CreatePhoto(convTokenID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// GetAll godoc
// @Summary Get all photo
// @Description This endpoint for get all photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.JSONSwaggerCreatedResponses
// @Router /photos [get]
// @Security BearerToken
func (c *photoController) GetAll(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	responses, err := c.usecase.GetAllPhoto(convTokenID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Update godoc
// @Summary Update photo
// @Description This endpoint for update photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param photoId path int true "PhotoID"
// @Param services body dtos.PhotoRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /photos/{photoId} [put]
// @Security BearerToken
func (c *photoController) Update(ctx echo.Context) error {
	request := dtos.PhotoRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.UpdatePhoto(convTokenID, convParamID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Delete godoc
// @Summary Delete photo
// @Description This endpoint for delete photo
// @Tags Photos
// @Accept  json
// @Produce  json
// @Param photoId path int true "PhotoID"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /photos/{photoId} [delete]
// @Security BearerToken
func (c *photoController) Delete(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		return err
	}

	responses, err := c.usecase.DeletePhoto(convTokenID, convParamID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
