// Package service 邮件插件业务层。
package service

import (
	"mall-admin/server/plugin/email/utils"
)

// EmailService 邮件 Service 实现。
type EmailService struct{}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func (e *EmailService) EmailTest() (err error) {
	subject := "test"
	body := "test"
	err = utils.EmailTest(subject, body)
	return err
}

//@author: [maplepie](https://github.com/maplepie)
//@function: EmailTest
//@description: 发送邮件测试
//@return: err error
//@params to string 	 收件人
//@params subject string   标题（主题）
//@params body  string 	 邮件内容

func (e *EmailService) SendEmail(to, subject, body string) (err error) {
	err = utils.Email(to, subject, body)
	return err
}
