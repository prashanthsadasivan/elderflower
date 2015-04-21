package controllers

import (
	"elderflower/controllers/socket_handlers"
	"elderflower/models"
	"elderflower/services/appconnections"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
)

func socketController(c *gin.Context) websocket.Handler {
	return func(ws *websocket.Conn) {
		// In order to select between websocket messages and subscription events, we
		// need to stuff websocket events into a channel.
		c.Request.ParseForm()
		qr_secret := c.Request.Form.Get("qr_secret")
		var appconnection *appconnections.AppConnection
		if qr_secret == "" {
			var qr_secret string
			qr_secret, appconnection = appconnections.GenerateForQR()
			if websocket.JSON.Send(ws, &models.QrCodeSecret{models.WebsocketModel{"qr_delivery"}, qr_secret}) != nil {
				return
			}
			log.Printf("qrwasempty\n")
		} else {
			appconnection = GorcAppConnection(qr_secret, c)
			log.Printf("gotorcreated\n")
		}
		messagesToSend := appconnection.Start(ws)
		log.Printf("socket appconn %+v\n", appconnection)

		// Now listen for new events from either the websocket or the chatroom.
		for {
			select {
			case event := <-appconnection.Received:
				log.Printf("ws, sent on websocket: %+v\n", event)
				if websocket.JSON.Send(ws, &event) != nil {
					// They disconnected.
					log.Printf("here")
					return
				}
			case msg, ok := <-messagesToSend:
				// If the channel is closed, they disconnected.
				log.Printf("ws sent a message")
				if !ok {
					return
				}
				socket_handlers.Router[msg.MessageType](msg, appconnection, ws)
			}
		}
	}
}

func HandleSocket(c *gin.Context) {
	websocket.Handler(socketController(c.Copy())).ServeHTTP(c.Writer, c.Request)
}
