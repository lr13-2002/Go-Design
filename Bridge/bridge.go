package bridge

type IMsgSender interface {
	Send(msg string) error
}

type EmailMsgSender struct {
	emails []string
}

func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}
