package uler_tangga_ws_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/go-uler-tangga-api/internal/model"
)

var (
	GlobalHub *Hub
)

func Initialize() {
	GlobalHub = newHub()
	go GlobalHub.run()
	go GlobalHub.runTicker()
}

func HandleWs(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.WithContext(c.Request.Context()).Error(err)
		return
	}

	client := &Client{
		hub:  GlobalHub,
		conn: conn,
		send: make(chan []byte, 256),
		identiy: model.Identity{
			ID:                    c.Request.URL.Query().Get("id"),
			Name:                  c.Request.URL.Query().Get("id"),
			RoomID:                c.Request.URL.Query().Get("room_id"),
			RoomPlayerIndex:       0,
			RoomPlayerIndexString: "0",
		},
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
