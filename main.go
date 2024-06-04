package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Define a route for the root URL
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
