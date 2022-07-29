package main

import (
	"fmt"
	"net/smtp"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const SubscribersFile = "subscribers.txt"

type EmailAddress struct {
	Email string `form:"email" binding:"required"`
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

func sendMessage(emailAddress string, binanceResponse BinanceResponse) {
	check(godotenv.Load())

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("PASSWORD")

	to := []string{emailAddress}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	fromMsg := fmt.Sprintf("From: <%s>\r\n", from)
	toMsg := fmt.Sprintf("To: <%s>\r\n", to[0])
	subject := "Subject: BTC to UAH rate\r\n"
	body := fmt.Sprintf("%s: %s", binanceResponse.Symbol, binanceResponse.Price)

	msg := fromMsg + toMsg + subject + "\r\n" + body

	check(smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg)))
}
