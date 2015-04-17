package controllers

import (
	"code.google.com/p/rsc/qr"
	"github.com/gin-gonic/gin"
)

func QR(c *gin.Context) {
	c.Request.ParseForm()
	hostname := c.Request.Form.Get("hostname")
	code, err := qr.Encode(hostname, qr.H)
	if err != nil {
		panic(err)
	}
	png := code.PNG()
	c.Data(200, "image/png", png)
}
