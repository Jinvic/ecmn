package models

import "encoding/json"

const (
	TopicUserCreated          = "user.created"
	TopicUserUpdated          = "user.updated"
	TopicUserDeleted          = "user.deleted"
	TopicEchoCreated          = "echo.created"
	TopicEchoUpdated          = "echo.updated"
	TopicEchoDeleted          = "echo.deleted"
	TopicCommentCreated       = "comment.created"
	TopicCommentStatusUpdated = "comment.status.updated"
	TopicCommentDeleted       = "comment.deleted"
	TopicResourceUploaded     = "resource.uploaded"
	TopicSystemBackup         = "system.backup"
	TopicSystemExport         = "system.export"
	TopicSystemBackupSchedule = "system.backup_schedule.updated"
	TopicInboxClear           = "inbox.clear"
	TopicEchoUpdateCheck      = "ech0.update.check"
)

var ValidTopics = []string{
	TopicUserCreated,
	TopicUserUpdated,
	TopicUserDeleted,
	TopicEchoCreated,
	TopicEchoUpdated,
	TopicEchoDeleted,
	TopicCommentCreated,
	TopicCommentStatusUpdated,
	TopicCommentDeleted,
	TopicResourceUploaded,
	TopicSystemBackup,
	TopicSystemExport,
	TopicSystemBackupSchedule,
	TopicInboxClear,
	TopicEchoUpdateCheck,
}

type WebhookPayload struct {
	Topic      string          `json:"topic"`
	EventName  string          `json:"event_name"`
	PayloadRaw json.RawMessage `json:"payload_raw"`
	Metadata   json.RawMessage `json:"metadata"`
	OccurredAt int64           `json:"occurred_at"`
}
