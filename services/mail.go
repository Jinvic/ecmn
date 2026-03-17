package services

import (
	"ecmn/config"
	"ecmn/models"
	"ecmn/pkg/client"
	"ecmn/pkg/mail"
	"fmt"
	"strings"
	"time"
)

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) SendCommentNotificationEmail(comment models.Comment) error {
	apiConfig := config.Get().API
	apiClient := client.New(apiConfig.BaseURL, apiConfig.Token, time.Duration(apiConfig.Timeout)*time.Second)

	var echo models.Echo
	if err := apiClient.GetResource("/api/echo/"+comment.EchoID, &echo); err != nil {
		return fmt.Errorf("failed to get echo: %w", err)
	}

	var systemSetting models.SystemSetting
	if err := apiClient.GetResource("/api/settings", &systemSetting); err != nil {
		return fmt.Errorf("failed to get system setting: %w", err)
	}

	smtpConfig := config.Get().SMTP
	smtpClient, err := mail.GetSMTPClient(smtpConfig.Host, smtpConfig.Port, smtpConfig.Username, smtpConfig.Password)
	if err != nil {
		return fmt.Errorf("failed to get SMTP client: %w", err)
	}

	serverName := systemSetting.ServerName
	if serverName == "" {
		serverName = "Ech0"
	}
	serverLogo := systemSetting.ServerLogo
	if serverLogo == "" {
		serverLogo = "/Ech0.svg"
	}
	serverURL := systemSetting.ServerURL
	if serverURL == "" {
		serverURL = "https://ech0.example.com"
	}
	serverLogo = fmt.Sprintf("%s%s", serverURL, serverLogo)
	data := mail.CommentNotificationEmailData{
		Title:     serverName,
		Logo:      serverLogo,
		Host:      serverURL,
		Poster:    echo.Username,
		Commenter: comment.Nickname,
		CommentAt: comment.CreatedAt,
		Content:   comment.Content,
		EchoID:    comment.EchoID,
	}
	// 生成邮件内容
	emailBody, err := mail.GenerateCommentNotificationEmail(data)
	if err != nil {
		return err
	}

	getDomain := func(email string) string {
		index := strings.LastIndex(email, "@")
		domain := strings.ToLower(email[index+1:])
		return domain
	}

	// 附加头部字段
	from := smtpConfig.From
	to := smtpConfig.To
	toHeader := strings.Join(to, ", ")
	subject := systemSetting.ServerName
	domain := getDomain(smtpConfig.From)
	email := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Date: "+time.Now().Format(time.RFC1123Z)+"\r\n"+
			"Message-ID: <"+time.Now().Format("20060102150405")+"@%s>\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=utf-8\r\n"+
			"\r\n"+
			"%s",
		from, toHeader, subject, domain, emailBody)

	// 发送邮件
	if err := smtpClient.SendMail(from, to, strings.NewReader(email)); err != nil {
		return err
	}

	return nil
}
