package main

import (
	"crypto/tls"
	"fmt"
	"github.com/i-coder-robot/email_demo/config"
	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", config.FROM)
	m.SetHeader("To", config.TO)
	m.SetHeader("Subject", "B站面向加薪学习")
	m.SetBody("text/plain", "欢喜-《Go语言极简一本通》")
	d := gomail.NewDialer("smtp.qq.com", 465, config.FROM, config.PASSWORD)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("发送结束")
}
