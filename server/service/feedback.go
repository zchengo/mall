package service

import (
	"crypto/tls"
	"fmt"
	"imall/global"
	"imall/models/web"

	"gopkg.in/gomail.v2"
)

type WebFeedBackService struct {
}

// 发送反馈的问题到QQ邮箱
func (f *WebFeedBackService) Send(param web.FeedbackParam) error {

	smtp := global.Config.Feedback.QqSmtp
	email := global.Config.Feedback.QqEmail
	secret := global.Config.Feedback.QqEmailSecret
	// QQ邮箱：SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163邮箱：SMTP 服务器地址：smtp.163.com（端口：25）
	m := gomail.NewMessage()
	m.SetHeader("From", email) // 发件人
	m.SetHeader("To", email)   // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	m.SetHeader("Cc", email)   // 抄送，可以多个
	m.SetHeader("Bcc", email)  // 暗送，可以多个
	m.SetHeader("Subject", "商城后台-问题反馈")      // 邮件主题
	m.SetBody("text/html", param.Content)
	d := gomail.NewDialer(smtp, 25, email, secret)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("qq mail send error : %s", err)
		return err
	}
	return nil
}
