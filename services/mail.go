package services

type MailService struct{}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) SendCommentNotificationEmail() error {
	// TODO: Implement the logic to send a comment notification email
	return nil
}
