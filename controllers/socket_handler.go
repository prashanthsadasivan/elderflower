package controllers

import (
	"elderflower/controllers/socket_handlers"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
)

func socketController(c *gin.Context) websocket.Handler {
	return func(ws *websocket.Conn) {
		// In order to select between websocket messages and subscription events, we
		// need to stuff websocket events into a channel.
		c.Request.ParseForm()
		num := c.Request.Form.Get("num")
		log.Printf("num: %s\n", num)
		appconnection := GorcAppConnection(num, c)
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
