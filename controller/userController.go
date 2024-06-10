package controller

import (
	"github.com/gin-gonic/gin"

	"schedule-maker/service"
)

func createUser(c *gin.Context) {
	user := service.createUser
	
}
