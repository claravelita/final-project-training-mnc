package api

import (
	"net/http"
	"os"
	"time"

	"github.com/claravelita/final-project-training-mnc/api/controller"
	"github.com/claravelita/final-project-training-mnc/api/handler"
	customMiddleware "github.com/claravelita/final-project-training-mnc/api/middleware"
	"github.com/claravelita/final-project-training-mnc/common/helper"
	_ "github.com/claravelita/final-project-training-mnc/docs"
	"github.com/claravelita/final-project-training-mnc/infrastructure"
	"github.com/claravelita/final-project-training-mnc/repository"
	photoUsecase "github.com/claravelita/final-project-training-mnc/usecase/photo"
	userUsecase "github.com/claravelita/final-project-training-mnc/usecase/user"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Server Struct
type Server struct {
	Route *echo.Echo
}

// NewServer : Create Server Instance
func NewServer(e *echo.Echo) *Server {
	return &Server{
		e,
	}
}

func (server *Server) InitializeServer() {
	server.Route.Use(customMiddleware.UseCorsMiddleware())
	customMiddleware.UseCustomLogger(server.Route)
	handler.UseCustomValidatorHandler(server.Route)

	initDB := infrastructure.NewGormDB()
	AuthHelper, _ := helper.NewAuthHelper(os.Getenv("SECRET_AUTH"))

	groupUser := server.Route.Group("/users")
	userRepo := repository.NewUserRepository(initDB)
	userUsecase := userUsecase.NewUserImplementation(userRepo, AuthHelper)
	userController := controller.NewUserController(userUsecase)
	userController.Route(groupUser)

	groupPhoto := server.Route.Group("/photos")
	photoRepo := repository.NewPhotoRepository(initDB)
	photoUsecase := photoUsecase.NewPhotoImplementation(photoRepo, userRepo, AuthHelper)
	photoController := controller.NewPhotoController(photoUsecase)
	photoController.Route(groupPhoto)

	serverConfiguration := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))
}
