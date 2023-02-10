package render

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Response(ctx context.Context, c *gin.Context, bodyPayload interface{}, errMsg string, status int) {
	success := true
	if status != 200 {
		success = false
	}

	// logrus.Infof("BODY RESPONSE: %+v", bodyPayload)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header(
		"Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Animapu-User-Uid, X-Visitor-Id, X-From-Path",
	)
	c.JSON(status, gin.H{
		"success":       success,
		"data":          bodyPayload,
		"error_message": errMsg,
	})
}
