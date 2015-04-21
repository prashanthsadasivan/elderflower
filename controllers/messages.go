package controllers

import (
	"elderflower/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Messages_Receive(c *gin.Context) {
	c.Request.ParseForm()
	qr_secret := c.Request.Form.Get("qr_secret")
	numFrom := c.Request.Form.Get("numFrom")
	messageReceived := c.Request.Form.Get("messageReceived")
	log.Printf("received text message %s %s", numFrom, messageReceived)
	sms := models.SMSMessage{}
	sms.MessageType = "SMS/Received"
	sms.Num = strings.TrimPrefix(numFrom, "+1")
	sms.Message = messageReceived
	appconnection := GorcAppConnection(qr_secret, c)
	log.Printf("appconn %+v\n", appconnection)
	go func() {
		log.Printf("beforeSend")
		appconnection.Received <- sms
		log.Printf("received, sent on channel\n")
	}()
	c.String(http.StatusAccepted, "")
}
