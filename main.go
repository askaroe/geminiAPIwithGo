package main

import (
	"github.com/askaroe/geminiApiGolang/controllers"
	"github.com/askaroe/geminiApiGolang/jsonlog"
	"github.com/askaroe/geminiApiGolang/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var logger *jsonlog.Logger

func main() {
	logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	err := godotenv.Load()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	service := services.NewService(logger)
	controller := controllers.NewController(service, logger)

	r := gin.Default()

	api := r.Group("api/v1")
	{
		api.GET("/healthcheck", controller.HealthCheck)
		api.POST("/aitest", controller.GetUserSummary)
	}

	logger.PrintInfo("server started", nil)

	err = r.Run()

	if err != nil {
		logger.PrintFatal(err, nil)
		return
	}

}
