# Ech0 Comment Mail Notifier (ecmn)

为 Ech0 项目提供评论通知的 Webhook 服务端，通过 SMTP 发送邮件通知。

## 快速开始

### 配置

编辑 `config.yaml`：

```yaml
server:
  port: 8080
  mode: release

webhook:
  secret: "your-webhook-secret"

logging:
  level: "info"
  format: "json"

smtp:
  host: "smtp.example.com"
  port: 587
  username: "noreply@example.com"
  password: "smtp-password"
  from: "noreply@example.com"
  to:
    - "admin@example.com"

api:
  base_url: "http://localhost:3000"
  token: "your-bearer-token"
  timeout: 30
```

### 运行

```bash
./ecmn.exe
```

## 项目结构

```
ecmn/
├── config/          # 配置加载
├── handlers/        # HTTP 处理器
├── logger/          # 日志封装
├── middleware/      # 中间件（签名验证）
├── models/          # 数据模型
├── pkg/
│   ├── client/      # HTTP 客户端
│   └── mail/        # 邮件发送
├── router/          # 路由配置
├── services/        # 业务逻辑
├── doc/             # 文档
├── main.go
└── config.yaml
```

## 构建

```bash
go build -o ecmn.exe .
```

## Docker

### Docker Compose

```bash
docker-compose up -d
```

### Docker Hub

镜像地址：`your-dockerhub-username/ecmn`

```bash
docker pull your-dockerhub-username/ecmn:latest
docker run -d -p 8080:8080 -v ./config.yaml:/app/config.yaml:ro your-dockerhub-username/ecmn
```
