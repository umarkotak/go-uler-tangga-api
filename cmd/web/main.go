package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/controller/uler_tangga_controller"
	"github.com/umarkotak/go-uler-tangga-api/internal/controller/uler_tangga_ws_controller"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
	"github.com/umarkotak/go-uler-tangga-api/internal/utils/render"
)

func main() {
	logrus.SetReportCaller(true)

	singleton.Initialize()
	uler_tangga_ws_controller.Initialize()

	r := gin.New()
	r.Use(CORSMiddleware())

	r.GET("/uler_tangga/start", uler_tangga_ws_controller.HandleWs)
	r.GET("/uler_tangga/room/list", uler_tangga_controller.GetRoomList)
	r.GET("/uler_tangga/room/:room_id/log", uler_tangga_controller.GetRoomMoveLog)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6001"
	}
	r.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			render.Response(c.Request.Context(), c, nil, "", 200)
			return
		}
		c.Next()
	}
}
