package socket_handlers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"golang.org/x/net/websocket"
)

type WebsocketMessageHandler func(sms models.SMSMessage, appconnection *appconnections.AppConnection, ws *websocket.Conn) bool

var Router map[string]WebsocketMessageHandler

func init() {
	Router = make(map[string]WebsocketMessageHandler)
	Router["ping"] = handlePing
	Router["SMS/Send"] = sendSmsMessage
}
