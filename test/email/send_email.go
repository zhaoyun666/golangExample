package email

import (
	"github.com/go-gomail/gomail"
	"log"
)

func send() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "pms@17car.com.cn" /*"发件人地址"*/, "发件人") // 发件人

	m.SetHeader("To",
		m.FormatAddress("Chuck.zhao@carlt.com.cn", "收件人")) // 收件人

	m.SetHeader("Subject", "liic测试")     // 主题

	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/html", "我是正文") // 正文

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "pms@17car.com.cn", "Grassroots1987") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}

	log.Println("done.发送成功")
}
