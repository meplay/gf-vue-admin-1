package email

type Request struct {
	To      string `json:"to"`      // 邮件发送给谁
	Body    string `json:"body"`    // 邮件内容
	Subject string `json:"subject"` // 邮件标题
}
