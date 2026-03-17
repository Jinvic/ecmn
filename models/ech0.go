package models

import "time"

// Comment model

type CommentStatus string

const (
	CommentStatusPending  CommentStatus = "pending"
	CommentStatusApproved CommentStatus = "approved"
	CommentStatusRejected CommentStatus = "rejected"
)

type CommentSourceType string

const (
	CommentSourceTypeGuest  CommentSourceType = "guest"
	CommentSourceTypeSystem CommentSourceType = "system"
)

type Comment struct {
	ID        string            `gorm:"type:char(36);primaryKey" json:"id"`
	EchoID    string            `gorm:"type:char(36);not null;index" json:"echo_id"`
	UserID    *string           `gorm:"type:char(36);index" json:"user_id,omitempty"`
	Nickname  string            `gorm:"size:100;not null;index" json:"nickname"`
	Email     string            `gorm:"size:255;not null;index" json:"email"`
	Website   string            `gorm:"size:255" json:"website,omitempty"`
	Content   string            `gorm:"type:text;not null" json:"content"`
	Status    CommentStatus     `gorm:"type:varchar(20);not null;index" json:"status"`
	Hot       bool              `gorm:"not null;default:false;index" json:"hot"`
	IPHash    string            `gorm:"size:128;index" json:"-"`
	UserAgent string            `gorm:"size:512" json:"-"`
	Source    CommentSourceType `gorm:"type:varchar(20);not null;index" json:"source"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// Echo model

// Echo 定义Echo实体
type Echo struct {
	ID        string         `gorm:"type:char(36);primaryKey"                      json:"id"`
	Content   string         `gorm:"type:text;not null"                            json:"content"`
	Username  string         `gorm:"type:varchar(100)"                             json:"username,omitempty"`
	EchoFiles []EchoFile     `gorm:"foreignKey:EchoID;constraint:OnDelete:CASCADE" json:"echo_files,omitempty"`
	Layout    string         `gorm:"type:varchar(50);default:'waterfall'"          json:"layout,omitempty"`
	Private   bool           `gorm:"default:false;index:idx_echos_private_created,priority:1" json:"private"`
	UserID    string         `gorm:"type:char(36);not null;index"                  json:"user_id"`
	Extension *EchoExtension `gorm:"foreignKey:EchoID;constraint:OnDelete:CASCADE" json:"extension,omitempty"`
	Tags      []Tag          `gorm:"many2many:echo_tags;"                          json:"tags,omitempty"`
	FavCount  int            `gorm:"default:0"                                     json:"fav_count"`
	CreatedAt time.Time      `gorm:"index:idx_echos_private_created,priority:2"    json:"created_at"`
}

type EchoExtension struct {
	ID        string                 `gorm:"type:char(36);primaryKey"      json:"id"`
	EchoID    string                 `gorm:"type:char(36);not null;uniqueIndex" json:"echo_id"`
	Type      string                 `gorm:"type:varchar(100);not null"    json:"type"`
	Payload   map[string]interface{} `gorm:"serializer:json;type:text;not null" json:"payload"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

// EchoFile links a File to an Echo with ordering support.
type EchoFile struct {
	ID        string `gorm:"type:char(36);primaryKey"                        json:"id"`
	EchoID    string `gorm:"type:char(36);uniqueIndex:idx_echo_file;not null" json:"echo_id"`
	FileID    string `gorm:"type:char(36);uniqueIndex:idx_echo_file;not null" json:"file_id"`
	File      File   `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE" json:"file,omitempty"`
	SortOrder int    `gorm:"default:0"                                   json:"sort_order"`
}

type File struct {
	ID string `gorm:"type:char(36);primaryKey" json:"id"`

	// 存储键（本地文件名或对象存储 object key）
	Key string `gorm:"type:varchar(500);not null;uniqueIndex:idx_file_route,priority:4" json:"key"`

	StorageType string `gorm:"type:varchar(20);not null;uniqueIndex:idx_file_route,priority:1" json:"storage_type"` // local|object|external
	Provider    string `gorm:"type:varchar(50);uniqueIndex:idx_file_route,priority:2" json:"provider,omitempty"`    // object 提供商，如 aws/r2/minio/external
	Bucket      string `gorm:"type:varchar(120);uniqueIndex:idx_file_route,priority:3" json:"bucket,omitempty"`     // local/external 可空

	URL         string `gorm:"type:text" json:"url"` // 前端直链快照
	Name        string `gorm:"type:varchar(255)" json:"name"`
	ContentType string `gorm:"type:varchar(100)" json:"content_type,omitempty"`
	Size        int64  `gorm:"default:0" json:"size"`
	Width       int    `gorm:"default:0" json:"width,omitempty"`
	Height      int    `gorm:"default:0" json:"height,omitempty"`

	Category  string    `gorm:"type:varchar(20);index" json:"category"` // image|video|audio|document|file
	UserID    string    `gorm:"type:char(36);index;not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
	ID         string    `gorm:"type:char(36);primaryKey"              json:"id"`
	Name       string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
	UsageCount int       `gorm:"default:0"                             json:"usage_count"`
	CreatedAt  time.Time `                                             json:"created_at"`
}

// Setting model

type SystemSetting struct {
	SiteTitle     string `json:"site_title"`     // 站点标题
	ServerLogo    string `json:"server_logo"`    // 服务器Logo
	ServerName    string `json:"server_name"`    // 服务器名称
	ServerURL     string `json:"server_url"`     // 服务器地址
	AllowRegister bool   `json:"allow_register"` // 是否允许注册'
	DefaultLocale string `json:"default_locale"` // 站点默认语言（如 zh-CN / en-US）
	ICPNumber     string `json:"ICP_number"`     // 备案号
	FooterContent string `json:"footer_content"` // 自定义页脚内容
	FooterLink    string `json:"footer_link"`    // 自定义页脚链接
	MetingAPI     string `json:"meting_api"`     // Meting API 地址
	CustomCSS     string `json:"custom_css"`     // 自定义 CSS
	CustomJS      string `json:"custom_js"`      // 自定义 JS
}
