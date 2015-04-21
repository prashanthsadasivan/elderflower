package main

import (
	"elderflower/config"
	"elderflower/controllers"
	"elderflower/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//Initialize middlewares
	r.Use(middleware.Middlewares()...)

	//Statics, Javascripts and templates
	r.Static("/public", "./compiled/public")
	r.Static("/jsx", "./compiled/jsx")
	r.LoadHTMLGlob("compiled/html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//Application routes
	r.POST("/users", controllers.User_Create)
	r.GET("/QR", controllers.QR)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/websocket", func(c *gin.Context) {
		controllers.HandleSocket(c)
	})
	r.POST("/messages/receive", controllers.Messages_Receive)

	r.Run(config.Get("port"))
}
