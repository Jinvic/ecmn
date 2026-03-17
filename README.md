# Ech0 Comment Mail Notifier (ecmn)

为 Ech0 项目提供评论通知的 Webhook 服务端，通过 SMTP 发送邮件通知。

## 快速开始

### Docker Compose

```bash
curl -L https://raw.githubusercontent.com/Jinvic/ecmn/refs/heads/master/docker-compose.yml -o docker-compose.yml
curl -L https://raw.githubusercontent.com/Jinvic/ecmn/refs/heads/master/config.yaml.example -o config.yaml
docker-compose up -d
```

### Docker

镜像地址：`jinvic/ecmn`

```bash
curl -L https://raw.githubusercontent.com/Jinvic/ecmn/refs/heads/master/config.yaml.example -o config.yaml
docker pull jinvic/ecmn:latest
docker run -d -p 8080:8080 -v ./config.yaml:/app/config.yaml:ro -e TZ=Asia/Shanghai jinvic/ecmn:latest
```

### 配置项

编辑 `config.yaml`：

```yaml
server:
  port: 8080              # 服务端口
  mode: release           # 运行模式 (debug/release)

webhook:
  secret: "your-webhook-secret"   # Webhook 签名密钥

logging:
  level: "info"           # 日志级别 (debug/info/warn/error)
  format: "json"          # 输出格式 (json/console)

smtp:
  host: "smtp.example.com"         # SMTP 服务器地址
  port: 587                        # SMTP 端口
  username: "noreply@example.com"  # SMTP 用户名
  password: "smtp-password"        # SMTP 密码
  from: "noreply@example.com"      # 发件人地址
  to:                              # 收件人列表
    - "admin@example.com"

api:
  base_url: "http://localhost:3000"  # Ech0 API 地址
  token: "your-bearer-token"         # API 访问 Token
  timeout: 30                        # 请求超时时间(秒)
```

## 项目结构

```bash
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
├── config.yaml
├── Dockerfile
└── docker-compose.yml
```
