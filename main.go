package main

import (
	"bufio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func returnBtcUahRate(context *gin.Context) {
	binanceResponse := getBtcUahRate()

	context.JSON(http.StatusOK, binanceResponse.Price)
}

func appendEmailToFile(context *gin.Context) {
	newEmailAddress := EmailAddress{}
	check(context.ShouldBindJSON(&newEmailAddress))

	if !isEmailValid(newEmailAddress.Email) {
		context.JSON(http.StatusBadRequest, "EMAIL_IS_NOT_VALID")
		return
	}

	file, err := os.OpenFile(SubscribersFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emailAddress := scanner.Text()
		if emailAddress == newEmailAddress.Email {
			context.JSON(http.StatusConflict, "EMAIL_SUBSCRIBED")
			return
		}
	}
	check(scanner.Err())

	file.Write([]byte(newEmailAddress.Email + "\n"))

	context.Status(http.StatusCreated)
}

func sendEmails(context *gin.Context) {
	binanceResponse := getBtcUahRate()

	if !isFileExists(SubscribersFile) {
		context.Status(http.StatusOK)
		return
	}

	file, err := os.Open(SubscribersFile)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emailAddress := scanner.Text()
		sendMessage(emailAddress, binanceResponse)
	}
	check(scanner.Err())

	context.Status(http.StatusOK)
}

func main() {
	r := gin.Default()
	r.GET("/rate", returnBtcUahRate)
	r.POST("/subscribe", appendEmailToFile)
	r.POST("/sendEmails", sendEmails)
	r.Run("0.0.0.0:8080")
}
