package app

import (
	"dating/app/config"
	"dating/constants"
	"dating/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	godotenv "github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

var (
	router *gin.Engine
)

func RunApp() {
	router = gin.Default()
	setupConfig()
	initEnv()

	repo := config.InitRepository()
	services := config.InitServices(*repo)

	controller := controllers.InitController(*services)

	controllerList := initRoutes(controller)
	controllerList.registerRoutes()

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("PORT_ENV"))); err != nil {
		panic(err)
	}
}

func setupConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	envVariable := godotenv.Load(environmentPath)
	if envVariable != nil {
		log.Fatal("Error loading .env file")
	}

}

func initEnv() {
	constants.Port = os.Getenv("PORT_ENV")
	constants.MySQLHost = os.Getenv("DB_HOST")
	constants.MySQLUser = os.Getenv("DB_USER")
	constants.MySQLPassword = os.Getenv("DB_PASSWORD")
	constants.MySQLPort = os.Getenv("DB_PORT")
	constants.MySQLDBName = os.Getenv("DB_NAME")
	constants.MySQLDBSchema = os.Getenv("DB_SCHEMA")
	constants.AutoMigrate = os.Getenv("AUTOMIGRATE")
}
