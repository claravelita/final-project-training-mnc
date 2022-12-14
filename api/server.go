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
	commentUsecase "github.com/claravelita/final-project-training-mnc/usecase/comment"
	photoUsecase "github.com/claravelita/final-project-training-mnc/usecase/photo"
	socialMediaUsecase "github.com/claravelita/final-project-training-mnc/usecase/socialMedia"
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

	groupComment := server.Route.Group("/comments")
	commentRepo := repository.NewCommentRepository(initDB)
	commentUsecase := commentUsecase.NewCommentImplementation(commentRepo, photoRepo, userRepo, AuthHelper)
	commentController := controller.NewCommentController(commentUsecase)
	commentController.Route(groupComment)

	groupSocialMedia := server.Route.Group("/socialmedias")
	socialMediaRepo := repository.NewSocialMediaRepository(initDB)
	socialMediaUsecase := socialMediaUsecase.NewSocialMediaImplementation(socialMediaRepo, commentRepo, photoRepo, userRepo, AuthHelper)
	socialMediaController := controller.NewSocialMediaController(socialMediaUsecase)
	socialMediaController.Route(groupSocialMedia)

	serverConfiguration := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	server.Route.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Final Project Training MNC x Hacktiv8 OK")
	})
	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)
	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))
}
