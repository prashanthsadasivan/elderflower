package models

import (
	"github.com/go-gorp/gorp"
	"time"
)

type WebsocketModel struct {
	MessageType string
}

type SMSMessage struct {
	WebsocketModel
	Num        string
	Message    string
	MessageId  int
	CreateDate time.Time
}

type SoleUser struct {
	Number string
	GcmId  string
	UserId int
}

func GetSoleUser(txn *gorp.Transaction) *SoleUser {
	soleUser := &SoleUser{}
	err := txn.SelectOne(soleUser, "select * from SoleUser")
	if err != nil {
		panic(err)
	}
	return soleUser
}
