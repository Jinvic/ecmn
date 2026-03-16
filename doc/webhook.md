# Webhook 使用说明

快速了解事件范围、请求结构、签名校验与失败处理策略。

## 支持事件（Event Topics）

以下 topic 会进入 Webhook 投递流程。

- `user.created`
- `user.updated`
- `user.deleted`
- `echo.created`
- `echo.updated`
- `echo.deleted`
- `comment.created`
- `comment.status.updated`
- `comment.deleted`
- `resource.uploaded`
- `system.backup`
- `system.export`
- `system.backup_schedule.updated`
- `inbox.clear`
- `ech0.update.check`

## 请求头（Headers）

- `X-Ech0-Event`：当前事件主题，例如 `echo.created`
- `X-Ech0-Event-ID`：事件唯一标识，建议用于幂等去重
- `X-Ech0-Timestamp`：Unix 秒级时间戳（UTC），建议用于防重放
- `X-Ech0-Signature`：可选签名头，格式 `sha256=<hex>`
- `User-Agent`：固定为 `Ech0-Webhook-Client`

## 请求体（Body）

- `topic`：事件主题
- `event_name`：事件类型名（后端事件结构名）
- `payload_raw`：事件原始业务数据
- `metadata`：附加元信息
- `occurred_at`：事件发生时间（UTC 时间戳）

## 请求示例

下面是一个典型的 Webhook JSON 请求体示例：

```json
{
  "topic": "echo.created",
  "event_name": "EchoCreatedEvent",
  "payload_raw": {
    "echo": {
      "id": "018f5e24-0fb7-7af0-a31b-a7ac0ad5e731",
      "content": "Hello from Ech0 webhook"
    },
    "user": {
      "id": "018f5e24-12f7-70e5-8a87-cf03a12bf10c",
      "username": "admin"
    }
  },
  "metadata": null,
  "occurred_at": 1710000000
}
```
