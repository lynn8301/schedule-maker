package User

import (
    "github.com/gin-gonic/gin"
	"net/http"
)

func userLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
        "status": "OK",
        "msg":    "FROM userLogin",
    })
}

func User(router *gin.Engine)  {
	userRouter := router.Group("/users")
	userRouter.GET("/login", userLogin)
}
