package bridge

type IMessageSender interface {
	Send(message string)
}

type MailSender struct{ mails []string }

func NewMailSender(mails []string) *MailSender { return &MailSender{mails: mails} }
func (s *MailSender) Send(message string)      { return }

type SmsSender struct{ phones []string }

func NewSmsSender(phones []string) *SmsSender { return &SmsSender{phones: phones} }
func (s *SmsSender) Send(message string)      { return }

type INotification interface {
	Notify(message string)
}

type SevereNotification struct {
	sender IMessageSender
}

func NewSevereNotification(sender IMessageSender) *SevereNotification {
	return &SevereNotification{
		sender: sender,
	}
}
func (n *SevereNotification) Notify(message string) {
	n.sender.Send(message)
}


