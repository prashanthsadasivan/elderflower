package socket_handlers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"encoding/json"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
	"strings"
)

func sendSmsMessage(message models.SMSMessage, appconn *appconnections.AppConnection, ws *websocket.Conn) bool {
	log.Printf("sendingMessage: %s\n", message.Message)

	client := http.Client{}

	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	datastring := string(payload[:])
	log.Printf("datastring: %s\n", datastring)
	log.Printf("regid: %s\n", appconn.RegId)
	log.Printf("gcm: %s\n", os.Getenv("GCM_AUTH_KEY"))

	request, err := http.NewRequest("POST", "https://android.googleapis.com/gcm/send", strings.NewReader("{\"registration_ids\":[\""+appconn.RegId+"\"], \"data\" : "+datastring+"}"))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "key="+os.Getenv("GCM_AUTH_KEY"))
	resp, err2 := client.Do(request)
	if err2 != nil {
		panic(err2)
	}
	if resp != nil {
		log.Printf("responseCode: %d\n", resp.StatusCode)
	}
	return true
}
