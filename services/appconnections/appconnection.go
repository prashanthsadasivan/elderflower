package appconnections

import (
	"elderflower/models"
	"golang.org/x/net/websocket"
	"log"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

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

func GetAppConnection(key string) *AppConnection {
	return conduits[key]
}

func New(regid, qr_secret string) *AppConnection {
	conduit := new(AppConnection)
	conduit.RegId = regid
	conduit.Received = make(chan models.SMSMessage)
	conduits[qr_secret] = conduit
	return conduit
}

func GenerateForQR() (string, *AppConnection) {
	conduit := new(AppConnection)
	conduit.Received = make(chan models.SMSMessage)
	key := randSeq(18)
	conduits[key] = conduit
	return key, conduit
}

func (c AppConnection) Start(ws *websocket.Conn) chan models.SMSMessage {
	messagesToSend := make(chan models.SMSMessage)
	go func() {
		var sms models.SMSMessage
		for {
			err := websocket.JSON.Receive(ws, &sms)
			if err != nil {
				log.Printf("err: %s\n", err.Error())
				close(messagesToSend)
				return
			}
			log.Printf("got message from ws: %+v\n", sms)
			messagesToSend <- sms
		}
	}()
	return messagesToSend
}
