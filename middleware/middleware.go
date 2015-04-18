package middleware

import (
	"github.com/gin-gonic/gin"
)

func Middlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		DatabaseMiddleware(),
	}
}
