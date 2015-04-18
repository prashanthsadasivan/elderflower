package controllers

import (
	"elderflower/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Receive(c *gin.Context) {
	c.Request.ParseForm()
	receiver := c.Request.Form.Get("receiver")
	numFrom := c.Request.Form.Get("numFrom")
	messageReceived := c.Request.Form.Get("messageReceived")
	fmt.Printf("received text message %s %s %s", receiver, numFrom, messageReceived)
	sms := models.SMSMessage{Num: strings.TrimPrefix(numFrom, "+1"), Message: messageReceived}
	appconnection := GorcAppConnection(receiver, c)
	fmt.Printf("appconn %+v\n", appconnection)
	go func() {
		fmt.Printf("beforeSend")
		appconnection.Received <- sms
		fmt.Printf("received, sent on channel\n")
	}()
	c.String(http.StatusAccepted, "")
}
