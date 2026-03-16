package router

import (
	"ecmn/handlers"
	"ecmn/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(whHandler *handlers.WebhookHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/webhook", middleware.SignatureVerify(), whHandler.HandleWebhook)

	return r
}
