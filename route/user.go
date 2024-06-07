package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "FROM userLogin",
	})
}

func UserRouters(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/login", userLogin)
	}
}
