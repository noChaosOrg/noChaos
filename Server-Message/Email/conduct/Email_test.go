package conduct

import (
	"github.com/noChaos1012/noChaos/Message/mail/defs"
	"testing"
)

func TestComments(t *testing.T) {
	t.Run("SenderMessage", testSenderMessage)
}

func testSenderMessage(t *testing.T) {
	mail := defs.Email{
		ServerHost: "smtp.163.com",
		ServerPort: 465,
		FromEmail:  "waschild.163.com",
		FromPwd:    "XJZDDLHZYACGJVWM",
		Toers:      []string{"497157441@qq.com", "1245838784@qq.com"},
		Subject:    "神秘邮件",
		Body:       "一封来自非繁的测试邮件",
	}
	err := SendMail(&mail)
	if err != nil {
		t.Errorf("Error of sendMessage：%v", err)
	}
}
