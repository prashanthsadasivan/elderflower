package controllers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"github.com/gin-gonic/gin"
	"log"
)

func User_Create(c *gin.Context) {
	c.Request.ParseForm()
	regId := c.Request.Form.Get("regId")
	num := c.Request.Form.Get("num")
	qr_secret := c.Request.Form.Get("qr_secret")
	log.Printf("regId: %s\n num:%s\n postnum: %s\n", regId, num, c.Request.PostFormValue("num"))
	appconnection := appconnections.GetAppConnection(qr_secret)
	appconnection.RegId = regId
	go func() {
		confirmationMessage := models.SMSMessage{}
		confirmationMessage.MessageType = "phone_confirmed"
		appconnection.Received <- confirmationMessage
	}()
	txn := Txn(c)
	results, err := txn.Select(models.SoleUser{}, "select * from SoleUser")
	if err != nil {
		panic(err)
	}
	log.Printf("results: %s\n", results)
	var newSoleUser *models.SoleUser
	if len(results) >= 1 {
		numDeleted, err := txn.Delete(results...)
		if err != nil {
			log.Printf("err: %s\n", err)
		}
		log.Printf("numDeleted: %d\n", numDeleted)
	}
	newSoleUser = &models.SoleUser{Number: num, GcmId: regId, QrSecret: qr_secret}
	insertErr := txn.Insert(newSoleUser)
	if insertErr != nil {
		log.Printf("here err: %s\n", insertErr)
	}

	c.JSON(201, newSoleUser)
}
