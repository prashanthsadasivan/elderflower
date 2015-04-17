package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smswebproxy/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Static("/public", "./compiled/public")
	r.LoadHTMLGlob("compiled/html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/QR", controllers.QR)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":2020")
}
