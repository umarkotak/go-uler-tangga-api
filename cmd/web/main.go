package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/controller/uler_tangga_controller"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
)

func main() {
	logrus.SetReportCaller(true)

	singleton.Initialize()
	uler_tangga_controller.Initialize()

	r := gin.New()
	r.GET("/uler_tangga/start", uler_tangga_controller.HandleWs)
	port := os.Getenv("PORT")
	if port == "" {
		port = "6001"
	}
	r.Run(":" + port)
}
