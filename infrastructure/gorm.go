package infrastructure

import (
	"fmt"
	"os"

	"github.com/claravelita/final-project-training-mnc/common"
	"github.com/claravelita/final-project-training-mnc/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func NewGormDB() *gorm.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	if dbInstance == nil {
		var err error

		DBInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			panic(err)
		}

		if common.IsMigrate && dbInstance == nil {
			err = DBInstance.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{})
			if err != nil {
				panic(err)
			}

			if err == nil {
				fmt.Println("Migrate Suceed!")
			}
		}

		dbInstance = DBInstance
		return dbInstance
	} else {
		return dbInstance
	}
}
