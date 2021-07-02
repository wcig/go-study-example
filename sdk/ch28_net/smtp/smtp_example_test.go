package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
	"testing"
)

var (
	username  = "user@example.com"                                          // 发件邮箱账号
	password  = "password"                                                  // 发件邮箱密码
	host      = "mail.example.com"                                          // 服务器域名
	addr      = "mail.example.com:25"                                       // 服务器地址（包括端口）
	from      = "sender@example.org"                                        // 发件人
	to        = []string{"recipient@example.net"}                           // 收件人
	cc        = "recipient@example.net"                                     // 抄送
	msgFormat = "To: %s\r\nFrom: %s\r\nCc: %s\r\nSubject: %s\r\n\r\n%s\r\n" // to,from,cc,subject,body
)

// 发件人和收件人设置昵称：from、to修改为格式：nickname<from/to>
func TestSendMail(t *testing.T) {
	auth := smtp.PlainAuth("", username, password, host)
	subject := "test subject"
	body := "test body"
	msg := []byte(fmt.Sprintf(msgFormat, strings.Join(to, ","), from, cc, subject, body))
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
