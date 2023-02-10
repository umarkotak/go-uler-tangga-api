package uler_tangga_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
	"github.com/umarkotak/go-uler-tangga-api/internal/singleton"
	"github.com/umarkotak/go-uler-tangga-api/internal/utils/render"
)

func GetRoomList(c *gin.Context) {
	rooms := []model.Room{}

	for _, room := range singleton.GetWorld().RoomMap {
		rooms = append(rooms, room)
	}

	render.Response(c.Request.Context(), c, rooms, "", 200)
}

func GetRoomMoveLog(c *gin.Context) {
	render.Response(
		c.Request.Context(),
		c,
		singleton.GetWorld().RoomMap[c.Param("room_id")].MoveLogs,
		"",
		200,
	)
}
