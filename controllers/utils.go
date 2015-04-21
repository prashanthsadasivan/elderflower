package controllers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

func GorcAppConnection(qr_secret string, c *gin.Context) *appconnections.AppConnection {
	appconnection := appconnections.GetAppConnection(qr_secret)
	if appconnection == nil {
		soleUser := models.GetSoleUser(Txn(c))
		appconnection = appconnections.New(soleUser.GcmId, soleUser.QrSecret)
	}

	return appconnection
}

func Txn(c *gin.Context) *gorp.Transaction {
	return c.MustGet("txn").(*gorp.Transaction)
}
