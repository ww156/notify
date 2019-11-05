/*
# @Author ww
# @Time 2019/10/11 9:19
# @File email.go
*/
package email

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"gopkg.in/gomail.v2"
	"net/mail"
	"notify/config"
	"notify/template"
	"path/filepath"
)

const (
	DefaultEmailSubject = "模板邮件"
)

var (
	DefaultTemplate = filepath.Join("..", "template", "email.tmpl")
)

// Email implements a Notifier for email notifications.
type Email struct {
	conf   *config.EmailConfig
	logger log.Logger
}

func New(c *config.EmailConfig, l log.Logger) *Email {
	return &Email{conf: c, logger: l}
}

func (e *Email) Notify(ctx context.Context) (bool, error) {
	var (
		err error
	)
	conf := e.conf
	// 发件人
	from_addrs, err := mail.ParseAddressList(conf.From)
	if err != nil {
		return false, fmt.Errorf("parse 'from' addresses\n%w", err)
	}
	if len(from_addrs) != 1 {
		return false, fmt.Errorf("must be exactly one 'from' address (got: %d)\n%w", len(from_addrs), err)
	}

	// 收件人
	to_addrs, err := mail.ParseAddressList(conf.To)
	if err != nil {
		return false, fmt.Errorf("parse 'to' addresses\n%w", err)
	}

	// 设置邮件头
	m := gomail.NewMessage()
	m.SetHeader("From", from_addrs[0].Address)
	m.SetHeader("To", parseEmilAddress(to_addrs)...)

	// 抄送人
	cc := e.conf.Cc
	if cc != "" {
		cc_addrs, err := mail.ParseAddressList(cc)
		if err != nil {
			return false, fmt.Errorf("parse 'cc' addresses\n%w", err)
		}
		m.SetHeader("Cc", parseEmilAddress(cc_addrs)...)
	}

	// 密送人
	bcc := e.conf.Bcc
	if bcc != "" {
		bcc_addrs, err := mail.ParseAddressList(bcc)
		if err != nil {
			return false, fmt.Errorf("parse 'bcc' addresses\n%w", err)
		}
		m.SetHeader("Bcc", parseEmilAddress(bcc_addrs)...)
	}

	// set the email Subject
	if conf.Subject == "" {
		conf.Subject = DefaultEmailSubject
	}
	m.SetHeader("Subject", conf.Subject)

	if conf.HTML != "" {
		html, err := template.ExecuteHTMLString(conf.HTML, conf.Data)
		if err != nil {
			return false, fmt.Errorf("parse html template\n%w", err)
		}
		m.SetBody("text/html", html)
	} else if conf.Text != "" {
		text, err := template.ExecuteHTMLString(conf.Text, conf.Data)
		if err != nil {
			return false, fmt.Errorf("parse text template\n%w", err)
		}
		m.SetBody("text/html", text)
	} else {
		html, err := template.ExecuteHTMLString(DefaultTemplate, map[string]string{"From": from_addrs[0].Address})
		if err != nil {
			return false, fmt.Errorf("parse html template\n%w", err)
		}
		m.SetBody("text/html", html)
	}

	d := gomail.NewDialer(conf.Smarthost.Host, conf.Smarthost.Port, conf.AuthUsername, string(conf.AuthPassword))

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}

	return true, nil
}

func parseEmilAddress(addresses []*mail.Address) []string {
	arr := []string{}
	for _, item := range addresses {
		arr = append(arr, item.String())
	}
	return arr
}
