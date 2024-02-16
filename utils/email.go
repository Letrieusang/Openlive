package utils

import (
	"api-openlive/common"
	"api-openlive/config"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

func GenRandomOTP() string {
	rangOtp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	otp := ""
	for i := 0; i < common.USER_OTP_LENG; i++ {
		rand.Seed(time.Now().UnixNano() + int64(len(rangOtp)))
		randomIndex := rand.Intn(len(rangOtp))
		otp += fmt.Sprint(rangOtp[randomIndex])
	}
	return otp
}

func SendMail(subject, otp string, to []string) error {
	// Sender data.
	from := config.GetConfig().GetString("mail.email")
	password := config.GetConfig().GetString("mail.password")
	log.Println(from, password)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, errFile := template.ParseFiles("./common/mailTmp.html")
	if errFile != nil {
		return errFile
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	t.Execute(&body, struct {
		Otp string
	}{
		Otp: otp,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func SendVerifyMail(UserName, verifyToken string, UserId int64, to []string) error {
	// Sender data.
	from := config.GetConfig().GetString("mail.email")
	password := config.GetConfig().GetString("mail.password")
	domain := config.GetConfig().GetString("server.change_domain")
	logo := config.GetConfig().GetString("server.logo")
	log.Println(from, password)

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, errFile := template.ParseFiles("./common/verifyMail.html")
	if errFile != nil {
		return errFile
	}
	subject := "Day la mail gui tu he thong opv market"
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	VerifyLink := fmt.Sprintf("%sverify-email/?verify_token=%s&UserId=%v\n", domain, verifyToken, UserId)
	t.Execute(&body, struct {
		Logo string
		Name string
		Web  string
	}{
		Logo: logo,
		Name: UserName,
		Web:  VerifyLink,
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}
