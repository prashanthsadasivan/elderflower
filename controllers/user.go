package controllers

import (
	"elderflower/models"
	"elderflower/services/appconnections"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Create(regId, num string, c *gin.Context) {
	fmt.Printf("regId: %s\n num:%s\n postnum: %s\n", regId, num, c.Request.PostFormValue("num"))
	appconnections.New(regId, num)
	txn := Txn(c)
	results, err := txn.Select(models.SoleUser{}, "select * from SoleUser")
	if err != nil {
		panic(err)
	}
	fmt.Printf("results: %s\n", results)
	if len(results) >= 1 {
		numDeleted, err := txn.Delete(results...)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
		fmt.Printf("numDeleted: %d\n", numDeleted)
	}
	newSoleUser := &models.SoleUser{Number: num, GcmId: regId}
	insertErr := txn.Insert(newSoleUser)
	if insertErr != nil {
		fmt.Printf("here err: %s\n", insertErr)
	}
	c.String(200, "created")
}
