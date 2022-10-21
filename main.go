package main

import (
	"fmt"
	"log"

	"github.com/claravelita/final-project-training-mnc/api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Final Project Training Golang MNC Service Api
// @version 1.0
// @description This is swagger for final project training golang mnc service with Hacktiv8
// @termsOfService http://swagger.io/terms/

// @contact.name Clara Velita Pranolo
// @contact.email claraavelitaa@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
func main() {
	log.Println("init main")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Println("application has started...")
	newEcho := echo.New()
	serverAPI := api.NewServer(newEcho)

	serverAPI.InitializeServer()
}
