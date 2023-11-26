package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	engine.Run(":8080")
}
