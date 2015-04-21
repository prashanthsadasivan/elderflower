package socket_handlers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"golang.org/x/net/websocket"
)

func handlePing(sms models.SMSMessage, appconnection *appconnections.AppConnection, ws *websocket.Conn) bool {
	return websocket.JSON.Send(ws, models.WebsocketModel{"pong"}) != nil
}
