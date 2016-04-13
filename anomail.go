package anomail

import (
	_ "fmt"
)

//content_type参数说明:
//1. text/plain  纯文本
//2. text/html  HTML文档
//3. text/xml  XML文档

func SendSimple(subject string, mail_from string, mail_to []string, content string, content_type ...string) error {
	Mail := New(subject, mail_from, mail_to)
	if 0 < len(content_type) {
		Mail.SetType(content_type[0])
	}
	err := Mail.Send(content)
	return err
}

func SendWithCc(subject string, mail_from string, mail_to []string, mail_cc []string, content string, content_type ...string) error {
	Mail := New(subject, mail_from, mail_to)
	Mail.SetCc(mail_cc)
	if 0 < len(content_type) {
		Mail.SetType(content_type[0])
	}
	err := Mail.Send(content)
	return err
}

func SendWithBcc(subject string, mail_from string, mail_to []string, mail_cc []string, mail_bcc []string, content string, content_type ...string) error {
	Mail := New(subject, mail_from, mail_to)
	Mail.SetCc(mail_cc)
	Mail.SetBcc(mail_bcc)
	if 0 < len(content_type) {
		Mail.SetType(content_type[0])
	}
	err := Mail.Send(content)
	return err
}
