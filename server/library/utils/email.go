package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/gogf/gf/frame/g"

	"github.com/jordan-wright/email"
)

func EmailTest(subject string, body string) error {
	to := strings.Split(g.Cfg("email").GetString("email.To"), ",")
	return send(to, subject, body)
}

// ErrorToEmail Error 发送邮件
func ErrorToEmail(subject string, body string) error {
	to := strings.Split(g.Cfg("email").GetString("email.To"), ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

func send(to []string, subject string, body string) error {
	from := g.Cfg("email").GetString("email.From")
	nickname := g.Cfg("email").GetString("email.Nickname")
	secret := g.Cfg("email").GetString("email.Secret")
	host := g.Cfg("email").GetString("email.Host")
	port := g.Cfg("email").GetInt("email.Port")
	isSSL := g.Cfg("email").GetBool("email.IsSSL")

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}
