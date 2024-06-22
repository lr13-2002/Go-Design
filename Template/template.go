package template

import "fmt"

type ISMS interface {
	send(content string, phone int) error
}

type sms struct {
	base ISMS
}

func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return nil
	}
	return s.base.send(content, phone)
}

type TelecomSms struct{}

func (TelecomSms) send(content string, phone int) error {
	fmt.Println("sned by telecom success")
	return nil
}

func NewTelecomSms() *sms {
	tel := &sms{base: TelecomSms{}}
	return tel
}
