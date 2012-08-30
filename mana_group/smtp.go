package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"time"
)

//
type MailUser struct {
	Identity string
	Username string
	Password string
	Host     string
	To       []string
	Port     string
}

var (
	charset = "Content-Type: text/plain; charset=UTF-8"
	subject = "Subject: mana warning"
	version = "MIME-Version: 1.0"
	date    = time.Now().Format(time.RFC1123Z)
)

func NewMailUser(p string) *MailUser {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		panic("read \"etc/mail\" failed")
	}
	var mu MailUser
	err = json.Unmarshal(b, &mu)
	if err != nil {
		panic("json.Unmarshal from \"etc/mail\" failed")
	}
	return &mu
}

func smtp_auth(mu *MailUser) smtp.Auth {
	return smtp.PlainAuth(mu.Identity, mu.Username, mu.Password, mu.Host)
}

//对于使用邮件过滤接收者的邮箱，需要设置To: <user@domain>,...
//同时添加From:<from@domain>
//Date: time.Now().Format(time.RFC1123Z) 
func warn_message(mu *MailUser, s string) []byte {
	var header = fmt.Sprintf("From: <%s>\r\n", mu.Username)
	if len(mu.To) > 0 {
		header += "To: "
		for i := 0; i < len(mu.To)-1; i++ {
			header += fmt.Sprintf("<%s>,", mu.To[i])
		}
		header += fmt.Sprintf("<%s>\r\n", mu.To[len(mu.To)-1])
	}
	header += fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s\r\n",
		subject, date, version, charset)
	if len(s) > 0 {
		msg := header + "\r\n" + s
		return []byte(msg)
	}
	return nil
}
