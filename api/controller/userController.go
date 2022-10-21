package controller

import (
	"fmt"
	"net/http"
	"strconv"

	authLibs "github.com/claravelita/final-project-training-mnc/api/middleware"
	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/claravelita/final-project-training-mnc/dtos"
	"github.com/claravelita/final-project-training-mnc/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userController struct {
	usecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) *userController {
	return &userController{usecase: usecase}
}

func (c *userController) Route(group *echo.Group) {
	group.POST("/register", c.Register)
	group.POST("/login", c.Login)
	group.PUT("/:userid", c.Update, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
	group.DELETE("", c.Delete, middleware.KeyAuth(authLibs.CheckJWTAndCheckAuthUser))
}

// Register godoc
// @Summary Register New User
// @Description This endpoint for create new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param services body dtos.UserRequest true "payload"
// @Success 201 {object} dtos.JSONSwaggerCreatedResponses
// @Router /users/register [post]
func (c *userController) Register(ctx echo.Context) error {
	request := dtos.UserRequest{}
	err := command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.RegisterUser(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Login godoc
// @Summary Login user
// @Description This endpoint for login user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param services body dtos.LoginRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /users/login [post]
func (c *userController) Login(ctx echo.Context) error {
	request := dtos.LoginRequest{}
	err := command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.LoginUser(request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Update godoc
// @Summary Update user
// @Description This endpoint for update user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param userid path int true "UserID"
// @Param services body dtos.UserUpdateRequest true "payload"
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /users/{userid} [put]
// @Security BearerToken
func (c *userController) Update(ctx echo.Context) error {
	request := dtos.UserUpdateRequest{}
	convParamID, err := strconv.Atoi(ctx.Param("userid"))
	if err != nil {
		return err
	}

	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	if convTokenID != convParamID {
		return ctx.JSON(http.StatusBadRequest, command.BadRequestResponses("Token or UserID incorrect"))
	}

	err = command.ValidateRequest(ctx, &request)
	if err != nil {
		return err
	}

	responses, err := c.usecase.UpdateUser(convParamID, request)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}

// Delete godoc
// @Summary Delete user
// @Description This endpoint for delete user
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} dtos.JSONSwaggerOKResponses
// @Router /users [delete]
// @Security BearerToken
func (c *userController) Delete(ctx echo.Context) error {
	convTokenID, err := strconv.Atoi(fmt.Sprint(ctx.Get("user_id")))
	if err != nil {
		return err
	}

	responses, err := c.usecase.DeleteUser(convTokenID)
	if err != nil {
		return err
	}
	return ctx.JSON(responses.Code, responses)
}
