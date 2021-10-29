package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

var Service = new(service)

type service struct{}

// Test 测试发送邮件
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *service) Test() error {
	subject := "test"
	body := "test"
	to := strings.Split(Config.To, ",")
	return s.send(to, subject, body)
}

// Send 发送邮件
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *service) Send(info *Request) error {
	to := strings.Split(info.To, ",")
	return s.send(to, info.Subject, info.Body)
}

// send 发送邮件 工具类
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *service) send(to []string, subject string, body string) error {
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	auth := smtp.PlainAuth("", Config.From, Config.Secret, Config.Host)
	e := email.NewEmail()
	if Config.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", Config.Nickname, Config.From)
	} else {
		e.From = Config.From
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", Config.Host, Config.Port)
	if Config.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: Config.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
