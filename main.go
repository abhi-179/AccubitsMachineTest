package main

import (
	logging "AccubitsMachineTest/accubits.machine.test/logger"
	"AccubitsMachineTest/accubits.machine.test/repository"
	"AccubitsMachineTest/accubits.machine.test/router"
	"AccubitsMachineTest/accubits.machine.test/service"
	"AccubitsMachineTest/docs"
	"log"

	config "AccubitsMachineTest/accubits.machine.test/config"
	database "AccubitsMachineTest/accubits.machine.test/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"sync"
)

var onceRest sync.Once

func main() {
	onceRest.Do(func() {
		e := echo.New()
		//Setting up the config
		config, err := config.GetConfig("./accubits.machine.test/env/")
		if err != nil {
			log.Print("config file not found")
		}
		//Setting up the Logger
		logger := logging.NewLogger(config.LogFile, config.LogLevel)
		logger.SetReportCaller(true)
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		docs.SwaggerInfo.Title = "Coursera Service API"
		docs.SwaggerInfo.Description = "Documentation coursera API v1.0"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:9000"
		docs.SwaggerInfo.Schemes = []string{"http"}
		e.GET("/swagger-ui/*", echoSwagger.WrapHandler)
		db := database.DB(config)
		testRepo := repository.NewtestRepository(db)
		testUc := service.NewtestUsecase(testRepo)
		router.NewTestController(e, testUc)
		if err := e.Start(config.HostPort); err != nil {
			logger.WithError(err).Fatal("not connected")
		}
	})
}
