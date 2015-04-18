package main

import (
	"elderflower/controllers"
	"elderflower/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//Initialize middlewares
	for _, middleware := range middleware.Middlewares() {
		r.Use(middleware)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Static("/public", "./compiled/public")
	r.Static("/jsx", "./compiled/jsx/app.js")
	r.LoadHTMLGlob("compiled/html/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/QR", controllers.QR)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":2020")
}
