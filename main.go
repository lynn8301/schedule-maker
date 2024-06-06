package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"user"
)

func main() {
	router := gin.Default()
	user.User(router)
	router.run(":8080")
}
