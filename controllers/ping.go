package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx , cancel := context.WithTimeout(context.Background(),time.Second)
		c.JSON(http.StatusOK, "Pong")
	}
}
