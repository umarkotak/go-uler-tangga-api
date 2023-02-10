package health_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umarkotak/go-uler-tangga-api/internal/utils/render"
)

func GetHealth(c *gin.Context) {
	render.Response(
		c.Request.Context(),
		c,
		map[string]interface{}{
			"health": "ok",
		},
		"",
		200,
	)
}
