package database

import (
	"fmt"
	"log"
	"os"

	"github.com/fvdime/go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectToDB ()  {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("failed to connect to database, \n", err)
		os.Exit(2)
	}

	log.Println("connected!")
	database.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations!")
	database.AutoMigrate(&models.Fact{})

	DB = DbInstance{
		Db: database,
	}

}

// dsn = data source name string