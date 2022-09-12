package provider

import (
	"errors"
	"fmt"
	"net/smtp"
	"time"

	structs "github.com/NothAmor/GoPusher/structs"
)

func Smtp(smtpAuth structs.SmtpRequestStruct) (structs.PusherResponse, error) {

	// 初始化返回值
	pusherReturn := structs.PusherResponse{
		Code:      500,
		Message:   "ERROR",
		Timestamp: time.Now().Unix(),
	}

	// 必填项不全，报错跳出函数
	if smtpAuth.Host == "" || smtpAuth.Account == "" || smtpAuth.Password == "" || smtpAuth.Port == 0 || smtpAuth.Sender == "" || len(smtpAuth.SendTo) == 0 || smtpAuth.Title == "" || smtpAuth.Content == "" || smtpAuth.MailType == "" {
		return pusherReturn, errors.New("缺失必填信息，请补全后重试！")
	}

	// SMTP认证
	auth := smtp.PlainAuth("", smtpAuth.Account, smtpAuth.Password, smtpAuth.Host)

	// 邮件格式初始化
	var contentType string = ""
	if smtpAuth.MailType == "html" {
		contentType = "Content-Type: text/html; charset=UTF-8"
	} else if smtpAuth.MailType == "plain" {
		contentType = "Content-Type: text/plain; charset=UTF-8"
	} else {
		return pusherReturn, errors.New("MailType值超出允许范围，允许值：{html, plain}")
	}

	// 遍历接收者数组，发送邮件
	for _, receiver := range smtpAuth.SendTo {
		mailContent := []byte(fmt.Sprintf("To: %s\r\nFrom: %s\r\nSubject: %s\r\n%s\r\n\r\n%s", receiver, smtpAuth.Sender, smtpAuth.Title, contentType, smtpAuth.Content))

		err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpAuth.Host, smtpAuth.Port), auth, smtpAuth.Sender, []string{receiver}, mailContent)
		if err != nil {
			return pusherReturn, fmt.Errorf("邮件发送错误！接收者：%s, 错误信息：%s, %s", receiver, err, fmt.Sprintf("%s:%d", smtpAuth.Host, smtpAuth.Port))
		}
	}

	pusherReturn.Code = 200
	pusherReturn.Message = "SUCCESS"

	return pusherReturn, nil
}
