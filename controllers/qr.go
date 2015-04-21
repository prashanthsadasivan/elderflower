package controllers

import (
	"code.google.com/p/rsc/qr"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func QR(c *gin.Context) {
	c.Request.ParseForm()
	hostname := c.Request.Form.Get("hostname")
	qr_secret := c.Request.Form.Get("qr_secret")
	log.Printf("qr secret: %s\n", qr_secret)
	mapEncoding := make(map[string]string)
	mapEncoding["hostname"] = hostname
	mapEncoding["qr_secret"] = qr_secret
	b, _ := json.Marshal(mapEncoding)
	code, err := qr.Encode(string(b), qr.H)
	if err != nil {
		panic(err)
	}
	png := code.PNG()
	c.Data(200, "image/png", png)
}
