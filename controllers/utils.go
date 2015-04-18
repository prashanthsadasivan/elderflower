package controllers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

func GorcAppConnection(num string, c *gin.Context) *appconnections.AppConnection {
	appconnection := appconnections.GetAppConnection(num)
	if appconnection == nil {
		soleUser := models.GetSoleUser(Txn(c))
		appconnection = appconnections.New(soleUser.GcmId, soleUser.Number)
	}

	return appconnection
}

func Txn(c *gin.Context) *gorp.Transaction {
	return c.MustGet("txn").(*gorp.Transaction)
}
