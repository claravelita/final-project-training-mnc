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

type commentController struct {
	usecase usecase.CommentUsecase
}

func NewCommentController(usecase usecase.CommentUsecase) *commentController {
	return &commentController{usecase: usecase}
}

func (c *commentController) Route(group *echo.Group) {
	group.POST("", c.Create, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.GET("", c.GetAll, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.PUT("/:commentId", c.Update, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.DELETE("/:commentId", c.Delete, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
}

// Create godoc
// @Summary Create photo
// @Description This endpoint for create photo
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param services body dtos.CommentRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerCreatedResponses
// @Router /comments [post]
// @Security BearerToken
func (c *commentController) Create(ctx echo.Context) error {
	request := dtos.CommentRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.CreateComment(convTokenID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// GetAll godoc
// @Summary Get all comment
// @Description This endpoint for get all comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /comments [get]
// @Security BearerToken
func (c *commentController) GetAll(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	responses, err := c.usecase.GetAllComment(convTokenID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Update godoc
// @Summary Update comment
// @Description This endpoint for update comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param commentId path int true "CommentID"
// @Param services body dtos.CommentUpdateRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /comments/{commentId} [put]
// @Security BearerToken
func (c *commentController) Update(ctx echo.Context) error {
	request := dtos.CommentUpdateRequest{}
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		return err
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.UpdateComment(convTokenID, convParamID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Delete godoc
// @Summary Delete comment
// @Description This endpoint for delete comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param commentId path int true "CommentID"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /comments/{commentId} [delete]
// @Security BearerToken
func (c *commentController) Delete(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	convParamID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		return err
	}

	responses, err := c.usecase.DeleteComment(convTokenID, convParamID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
