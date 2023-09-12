package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resp": "pong, Golang!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", returnPong)

	router.Run("localhost:3000")
}
