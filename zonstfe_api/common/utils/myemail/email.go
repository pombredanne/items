package myemail

import (
	"gopkg.in/gomail.v2"
	"time"
	"fmt"
)

func SendEmail(form, to []string, title, err string) {
	m := gomail.NewMessage()
	m.SetHeader("From", form...)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", fmt.Sprintf("%s %v", title, time.Now().Format("2006-01-02 15:04:05")))
	m.SetBody("text/html", err)
	//m.Attach("/home/Alex/lolcat.jpg")
	d := gomail.NewDialer("smtp.qq.com", 587, "1020300659@qq.com", "tlnptbwvxkhnbebj")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
