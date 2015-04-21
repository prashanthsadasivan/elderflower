package controllers

import (
	"elderflower/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Receive(c *gin.Context) {
	c.Request.ParseForm()
	receiver := c.Request.Form.Get("receiver")
	numFrom := c.Request.Form.Get("numFrom")
	messageReceived := c.Request.Form.Get("messageReceived")
	log.Printf("received text message %s %s %s", receiver, numFrom, messageReceived)
	sms := models.SMSMessage{Num: strings.TrimPrefix(numFrom, "+1"), Message: messageReceived}
	appconnection := GorcAppConnection(receiver, c)
	log.Printf("appconn %+v\n", appconnection)
	go func() {
		log.Printf("beforeSend")
		appconnection.Received <- sms
		log.Printf("received, sent on channel\n")
	}()
	c.String(http.StatusAccepted, "")
}
