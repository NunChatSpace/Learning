package Interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyInterceptor(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"result": "You fucked up"})
		c.Abort()
	}
}
