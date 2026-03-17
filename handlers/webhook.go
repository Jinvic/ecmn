package handlers

import (
	"encoding/json"
	"net/http"

	"ecmn/logger"
	"ecmn/models"
	"ecmn/services"

	"github.com/gin-gonic/gin"
)

const (
	HeaderEvent     = "X-Ech0-Event"
	HeaderEventID   = "X-Ech0-Event-ID"
	HeaderTimestamp = "X-Ech0-Timestamp"
)

type WebhookHandler struct {
	mailService *services.MailService
}

func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		mailService: services.NewMailService(),
	}
}

func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
	var payload models.WebhookPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error("Failed to parse webhook payload", logger.Err(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	eventID := c.GetHeader(HeaderEventID)
	timestamp := c.GetHeader(HeaderTimestamp)
	event := c.GetHeader(HeaderEvent)

	logger.Info("Received webhook",
		logger.String("event", event),
		logger.String("event_id", eventID),
		logger.String("timestamp", timestamp),
		logger.String("topic", payload.Topic),
		logger.String("event_name", payload.EventName))

	switch payload.Topic {
	case models.TopicCommentCreated:
		go func() {
			type PayloadRaw struct {
				Comment models.Comment
			}
			var payloadRaw PayloadRaw
			if err := json.Unmarshal(payload.PayloadRaw, &payloadRaw); err != nil {
				logger.Error("Failed to unmarshal comment", logger.Err(err))
				return
			}
			if err := h.mailService.SendCommentNotificationEmail(payloadRaw.Comment); err != nil {
				logger.Error("Failed to send comment notification email", logger.Err(err))
			}
			logger.Info("Comment notification email sent", logger.String("comment_id", payloadRaw.Comment.ID))
		}()
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "webhook received",
		"topic":      payload.Topic,
		"event_name": payload.EventName,
	})
}
