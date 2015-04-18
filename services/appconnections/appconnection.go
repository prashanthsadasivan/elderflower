package appconnections

import (
	"elderflower/models"
	"fmt"
	"golang.org/x/net/websocket"
)

var (
	conduits map[string]*AppConnection
)

func init() {
	conduits = make(map[string]*AppConnection)
}

type AppConnection struct {
	Received chan models.SMSMessage
	RegId    string
}

func GetAppConnection(number string) *AppConnection {
	return conduits[number]
}

func AddAppConnection(num string, conduit *AppConnection) {
	conduits[num] = conduit
}

func New(regid, num string) *AppConnection {
	conduit := new(AppConnection)
	conduit.RegId = regid
	conduit.Received = make(chan models.SMSMessage)
	conduits[num] = conduit
	return conduit
}

func (c AppConnection) Start(ws *websocket.Conn) chan models.SMSMessage {
	messagesToSend := make(chan models.SMSMessage)
	go func() {
		var sms models.SMSMessage
		for {
			err := websocket.JSON.Receive(ws, &sms)
			if err != nil {
				fmt.Printf("err: %s\n", err.Error())
				close(messagesToSend)
				return
			}
			fmt.Printf("got message from ws: %+v\n", sms)
			messagesToSend <- sms
		}
	}()
	return messagesToSend
}
